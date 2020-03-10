//go:generate go run github.com/99designs/gqlgen
package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	"github.com/asciiu/appa/lib/db"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/tinrab/retry"
)

type contextKey string

const (
	userContextKey = contextKey("user")
)

type graphQLServer struct {
	DB          *sql.DB
	redisClient *redis.Client
	//messageChannels map[string]chan *graph.Message
	userChannels map[string]chan string
	mutex        sync.Mutex
}

type Config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	DBURL    string `envconfig:"DB_URL"`
}

func NewGraphQLServer(config Config) (*graphQLServer, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	database, _ := db.NewDB(config.DBURL)

	retry.ForeverSleep(2*time.Second, func(_ int) error {
		_, err := client.Ping().Result()
		return err
	})
	return &graphQLServer{
		redisClient: client,
		DB:          database,
		//messageChannels: map[string]chan *graph.Message{},
		userChannels: map[string]chan string{},
		mutex:        sync.Mutex{},
	}, nil
}

func (s *graphQLServer) Serve(route string, port int) error {
	mux := http.NewServeMux()
	mux.Handle(
		route,
		handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: s}),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)
	mux.Handle("/playground", handler.Playground("GraphQL", route))

	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}

//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	graph "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	"github.com/asciiu/appa/lib/db"
	user "github.com/asciiu/appa/lib/user/models"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/tinrab/retry"
)

type contextKey string

const (
	// A private key for context that only this package can access. This is important
	// to prevent collisions between different context uses
	userContextKey = contextKey("user")
)

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *user.User {
	raw, _ := ctx.Value(userContextKey).(*user.User)
	return raw
}

type graphQLServer struct {
	DB              *sql.DB
	redisClient     *redis.Client
	messageChannels map[string]chan *graph.Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

type Config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	DBURL    string `envconfig:"APPA_API_GRAPHQL_DB_URL"`
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
		redisClient:     client,
		DB:              database,
		messageChannels: map[string]chan *graph.Message{},
		userChannels:    map[string]chan string{},
		mutex:           sync.Mutex{},
	}, nil
}

func (srv *graphQLServer) Mutation() generated.MutationResolver {
	return srv
}

func (srv *graphQLServer) Query() generated.QueryResolver {
	return srv
}

func makeWebsocketInitFunc(db *sql.DB) transport.WebsocketInitFunc {
	return func(ctx context.Context, initPayload handler.InitPayload) (context.Context, error) {
		tokenPayload := initPayload["token"]
		if tokenPayload == nil {
			return ctx, errors.New("Token Payload is nil")
		}
		token := tokenPayload.(string)
		fmt.Println(token)
		//user, err := auth.ParseTokenIntoUser(db, token)
		//if err != nil {
		//	return ctx, err
		//}
		return ctx, nil
	}
}

func (srv *graphQLServer) Serve(route string, port int) error {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Authorization", "API-Key", "Refresh", "Origin", "Accept"},
		AllowedMethods:     []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "TRACE", "PATCH"},
		ExposedHeaders:     []string{"set-authorization", "set-refresh"},
		OptionsPassthrough: true,
		Debug:              false,
	})

	cfg := generated.Config{Resolvers: srv}
	gqlsrv := handler.GraphQL(generated.NewExecutableSchema(cfg),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}),
	)
	//gqlsrv.AddTransport(transport.Options{})
	//gqlsrv.AddTransport(transport.GET{})
	//gqlsrv.AddTransport(transport.POST{})
	//gqlsrv.AddTransport(transport.Websocket{
	//	KeepAlivePingInterval: 10 * time.Second,
	//	Upgrader: websocket.Upgrader{
	//		CheckOrigin: func(r *http.Request) bool {
	//			return true
	//		},
	//	},
	//	InitFunc: transport.WebsocketInitFunc(makeWebsocketInitFunc(srv.DB)),
	//})

	router := chi.NewRouter()
	router.Use(authenticated(srv.DB))
	router.Use(corsConfig.Handler)
	router.Handle(route, gqlsrv)
	router.Handle("/playground", handler.Playground("GraphQL", route))
	fmt.Println(fmt.Sprintf("connect to http://localhost:%d/playground for GraphQL playground", port))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

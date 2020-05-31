//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	//handler2 "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"

	//"github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	graph "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	"github.com/asciiu/appa/lib/db/gopg"
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
	db              *pg.DB
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

	database, err := gopg.NewDB(config.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	retry.ForeverSleep(2*time.Second, func(_ int) error {
		_, err := client.Ping().Result()
		return err
	})
	return &graphQLServer{
		redisClient:     client,
		db:              database,
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

func (srv *graphQLServer) Subscription() generated.SubscriptionResolver {
	return srv
}

// func makeWebsocketInitFunc(db *sql.DB) transport.WebsocketInitFunc {
// 	return func(ctx context.Context, initPayload handler.InitPayload) (context.Context, error) {
// 		tokenPayload := initPayload["token"]
// 		if tokenPayload == nil {
// 			return ctx, errors.New("Token Payload is nil")
// 		}
// 		token := tokenPayload.(string)
// 		fmt.Println(token)
// 		//user, err := auth.ParseTokenIntoUser(db, token)
// 		//if err != nil {
// 		//	return ctx, err
// 		//}
// 		return ctx, nil
// 	}
// }

func getMediaBase(mId int) string {
	mediaRoot := "assets/media"
	return fmt.Sprintf("%s", mediaRoot)
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	log.Println(mediaFile)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
}

func streamHandler(response http.ResponseWriter, request *http.Request) {
	mid := chi.URLParam(request, "mId")

	mID, err := strconv.Atoi(mid)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	segName := chi.URLParam(request, "segName")
	//segName, ok := vars["segName"]
	if segName == "" {
		mediaBase := getMediaBase(mID)
		m3u8Name := "index.m3u8"
		serveHlsM3u8(response, request, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase(mID)
		serveHlsTs(response, request, mediaBase, segName)
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
	gqlsrv := handler.New(generated.NewExecutableSchema(cfg))
	gqlsrv.AddTransport(transport.Options{})
	gqlsrv.AddTransport(transport.GET{})
	gqlsrv.AddTransport(transport.POST{})
	gqlsrv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		//	InitFunc: transport.WebsocketInitFunc(makeWebsocketInitFunc(srv.DB)),
	})

	router := chi.NewRouter()
	//router.Use(authenticated(srv.DB))
	router.Use(corsConfig.Handler)

	router.Get("/media/{mId:[0-9]+}/stream/", streamHandler)
	router.Get("/media/{mId:[0-9]+}/stream/{segName:index[0-9]+.ts}", streamHandler)

	router.Handle(route, gqlsrv)
	router.Handle("/playground", playground.Handler("GraphQL", route))
	fmt.Println(fmt.Sprintf("connect to http://localhost:%d/playground for GraphQL playground", port))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

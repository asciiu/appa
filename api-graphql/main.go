package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/api-graphql/auth"
	gql "github.com/asciiu/appa/api-graphql/graphql"
	"github.com/asciiu/appa/lib/db"
	util "github.com/asciiu/appa/lib/util"
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	protoStory "github.com/asciiu/appa/story-service/proto/story"

	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/gqlerror"
)

const defaultPort = "8080"

// clean up stage refresh tokens in DB every 30 minutes
const cleanUpInterval = 30 * time.Minute

// routine to clean up refresh tokens in DB
func cleanDatabase(db *sql.DB) {
	for {
		time.Sleep(cleanUpInterval)
		error := tokenRepo.DeleteStaleTokens(db, time.Now())
		if error != nil {
			log.Error(error)
		}
	}
}

func main() {
	util.HelloThere("satori!")

	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	database, _ := db.NewDB(dbURL)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	service := k8s.NewService(micro.Name("graphql"))
	service.Init()

	resolver := gql.Resolver{
		DB:          database,
		StoryClient: protoStory.NewStoryService("stories", service.Client()),
	}

	router := chi.NewRouter()
	router.Use(auth.Secure(database))

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Refresh"},
		ExposedHeaders:   []string{"set-authorization", "set-refresh"},
		Debug:            false,
	}).Handler)

	router.Handle("/", handler.Playground("gql", "/graphql"))
	router.Handle("/graphql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolver}),
		handler.ErrorPresenter(
			func(ctx context.Context, e error) *gqlerror.Error {
				return graphql.DefaultErrorPresenter(ctx, e)
			},
		),
	))

	go cleanDatabase(database)

	//go func() {
	//	if err := service.Run(); err != nil {
	//		log.Error(fmt.Sprintf("nope! %s", err))
	//	}
	//}()

	log.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", defaultPort))
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}

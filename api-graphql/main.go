package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/asciiu/appa/api-graphql/auth"
	"github.com/asciiu/appa/api-graphql/graph"
	"github.com/asciiu/appa/lib/db"
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	util "github.com/asciiu/appa/lib/util"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
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

	_ = godotenv.Load("env/local.env")

	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	fmt.Println(dbURL)

	database, _ := db.NewDB(dbURL)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	//service := k8s.NewService(micro.Name("graphql"))
	//service.Init()

	resolver := &graph.Resolver{
		DB: database,
		//StoryClient: protoStory.NewStoryService("stories", service.Client()),
	}

	router := chi.NewRouter()
	router.Use(auth.Secure(database))

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Refresh"},
		ExposedHeaders:   []string{"set-authorization", "set-refresh"},
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	go cleanDatabase(database)

	//go func() {
	//	if err := service.Run(); err != nil {
	//		log.Error(fmt.Sprintf("nope! %s", err))
	//	}
	//}()

	log.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", defaultPort))
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}

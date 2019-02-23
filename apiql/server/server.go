package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/apiql"
	"github.com/asciiu/appa/apiql/auth"
	"github.com/asciiu/appa/common/db"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()

	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	database, _ := db.NewDB(dbURL)
	router.Use(auth.Middleware(database))
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Habibi", "/query"))
	router.Handle("/query", handler.GraphQL(apiql.NewExecutableSchema(apiql.Config{Resolvers: &apiql.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}

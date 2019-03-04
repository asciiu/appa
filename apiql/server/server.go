package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/apiql"
	"github.com/asciiu/appa/apiql/auth"
	repoUser "github.com/asciiu/appa/apiql/db/sql"
	"github.com/asciiu/appa/common/db"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

// clean up stage refresh tokens in DB every 30 minutes
const cleanUpInterval = 30 * time.Minute

// routine to clean up refresh tokens in DB
func cleanDatabase(db *sql.DB) {
	for {
		time.Sleep(cleanUpInterval)
		error := repoUser.DeleteStaleTokens(db, time.Now())
		if error != nil {
			log.Fatal(error)
		}
	}
}

func main() {
	router := chi.NewRouter()

	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	database, _ := db.NewDB(dbURL)
	resolver := apiql.Resolver{DB: database}
	router.Use(auth.Secure(database))

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Refresh"},
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Habibi", "/query"))
	router.Handle("/query", handler.GraphQL(apiql.NewExecutableSchema(apiql.Config{Resolvers: &resolver})))

	go cleanDatabase(database)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
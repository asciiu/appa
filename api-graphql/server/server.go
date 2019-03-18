package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/asciiu/appa/apiql"
	"github.com/asciiu/appa/apiql/auth"
	repoUser "github.com/asciiu/appa/apiql/db/sql"
	"github.com/asciiu/appa/common/db"
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
		ExposedHeaders:   []string{"set-authorization", "set-refresh"},
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Habibi", "/query"))
	router.Handle("/query", handler.GraphQL(apiql.NewExecutableSchema(apiql.Config{Resolvers: &resolver}),
		handler.ErrorPresenter(
			func(ctx context.Context, e error) *gqlerror.Error {
				return graphql.DefaultErrorPresenter(ctx, e)
			},
		),
	))

	go cleanDatabase(database)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}

package main

import (
	"log"

	"github.com/asciiu/appa/api-graphql-rocket/server"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// TODO read env from command arg
	_ = godotenv.Load("config/dev.env")

	var cfg server.Config
	err := envconfig.Process("", &cfg)
	check(err)

	srv, err := server.NewGraphQLServer(cfg)
	check(err)

	err = srv.Serve("/graphql", 8080)
	check(err)
}

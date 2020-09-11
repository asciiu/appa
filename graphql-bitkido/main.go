package main

import (
	"log"
	"os"

	"github.com/asciiu/appa/api-graphql-rocket/server"
	"github.com/asciiu/appa/lib/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("command line env file not found in command args")
	}

	envfile := argsWithoutProg[0]
	config.LoadEnv(envfile)

	var cfg server.Config
	err := envconfig.Process("", &cfg)
	check(err)

	srv, err := server.NewGraphQLServer(cfg)
	check(err)

	err = srv.Serve("/graphql", 8080)
	check(err)
}

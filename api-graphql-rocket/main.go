package main

import (
	"log"

	"github.com/asciiu/appa/api-graphql-rocket/server"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg server.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewGraphQLServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve("/graphql", 8080)
	if err != nil {
		log.Fatal(err)
	}
}

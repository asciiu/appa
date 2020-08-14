package main

import (
	"fmt"
	"os"

	"github.com/asciiu/appa/lib/config"
	"github.com/asciiu/appa/sandbox-01/grin"
	"github.com/blockcypher/libgrin/owner"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

func checkErr(label string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", label, err.Error())
		os.Exit(1)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("command line env file not found in command args")
	}

	envfile := argsWithoutProg[0]
	config.LoadEnv(envfile)

	var cfg grin.GrinConfig
	err := envconfig.Process("", &cfg)
	checkErr("process config", err)

	api := owner.NewSecureOwnerAPI(cfg.URL)

	err = api.Init()
	checkErr("failed to init api", err)

	err = api.Open(nil, "I am a warrior")
	checkErr("failed to open wallet", err)

	log.Infof("success: %s", api)
}

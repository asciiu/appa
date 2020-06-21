package main

import (
	"fmt"
	"os"

	"github.com/asciiu/appa/lib/config"
	"github.com/asciiu/appa/sandbox-01/grin"
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

	key, err := grin.InitSecureApi(cfg)
	checkErr("init secure api", err)

	pass := "I am a warrior"
	grin.OpenWallet(cfg, key, nil, &pass)
}

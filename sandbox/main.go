package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/lib/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/ybbus/jsonrpc"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type WalletInfo struct {
	AmountAwaitingConfirmation int `json:"amount_awaiting_confirmation"`
	AmountAwaitingFinalization int `json:"amount_awaiting_finalization"`
	AmountCurrentSpendable     int `json:"amount_currently_spendable"`
	LastConfirmedHeight        int `json:"last_confirmed_height"`
	MinConfirmations           int `json:"minimum_confirmations"`
}

type GrinConfig struct {
	URL string `envconfig:"GRIN_API_URL" required:"true"`
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("command line env file not found in command args")
	}

	envfile := argsWithoutProg[0]
	config.LoadEnv(envfile)

	var cfg GrinConfig
	err := envconfig.Process("", &cfg)
	check(err)

	rpcClient := jsonrpc.NewClient(cfg.URL)
	response, err := rpcClient.Call("retrieve_summary_info", true, 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}

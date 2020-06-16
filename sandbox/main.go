package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

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
	AmountAwaitingConfirmation int64 `json:"amount_awaiting_confirmation"`
	AmountAwaitingFinalization int64 `json:"amount_awaiting_finalization"`
	AmountCurrentSpendable     int64 `json:"amount_currently_spendable"`
	AmountImmature             int64 `json:"amount_immature"`
	AmountLocked               int64 `json:"amount_locked"`
	LastConfirmedHeight        int64 `json:"last_confirmed_height"`
	MinimumConfirmations       int64 `json:"minimum_confirmations"`
	Total                      int64 `json:"total"`
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

	resMap := make(map[string]interface{}, 0)
	response.GetObject(&resMap)

	if val, ok := resMap["Ok"]; ok {
		tup := val.([]interface{})
		wMap := tup[1].(map[string]interface{})

		amountAwaitingConfirmation, _ := strconv.ParseInt(wMap["amount_awaiting_confirmation"].(string), 10, 64)
		amountAwaitingFinalization, _ := strconv.ParseInt(wMap["amount_awaiting_finalization"].(string), 10, 64)
		amountCurrentSpendable, _ := strconv.ParseInt(wMap["amount_currently_spendable"].(string), 10, 64)
		amountImmature, _ := strconv.ParseInt(wMap["amount_immature"].(string), 10, 64)
		amountLocked, _ := strconv.ParseInt(wMap["amount_locked"].(string), 10, 64)
		lastConfirmedHeight, _ := strconv.ParseInt(wMap["last_confirmed_height"].(string), 10, 64)
		minimumConfirmations, _ := strconv.ParseInt(wMap["minimum_confirmations"].(string), 10, 64)
		total, _ := strconv.ParseInt(wMap["total"].(string), 10, 64)

		walletInfo := WalletInfo{
			AmountAwaitingConfirmation: amountAwaitingConfirmation,
			AmountAwaitingFinalization: amountAwaitingFinalization,
			AmountCurrentSpendable:     amountCurrentSpendable,
			AmountImmature:             amountImmature,
			AmountLocked:               amountLocked,
			LastConfirmedHeight:        lastConfirmedHeight,
			MinimumConfirmations:       minimumConfirmations,
			Total:                      total,
		}

		jsn, _ := json.Marshal(walletInfo)
		fmt.Printf("%s\n", jsn)
	}
}

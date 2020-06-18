package main

import (
	"encoding/json"
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

type OkJson struct {
	Ok []json.RawMessage
}

type Details struct {
	AmountAwaitingConfirmation string `json:"amount_awaiting_confirmation"`
	AmountAwaitingFinalization string `json:"amount_awaiting_finalization"`
	AmountCurrentlySpendable   string `json:"amount_currently_spendable"`
	AmountImmature             string `json:"amount_immature"`
	AmountLocked               string `json:"amount_locked"`
	LastConfirmedHeight        string `json:"last_confirmed_height"`
	MinimumConfirmations       string `json:"minimum_confirmations"`
	Total                      string `json:"total"`
}

func RPCSummary(conf GrinConfig) Details {
	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("retrieve_summary_info", true, 10)
	//printResult(*response, err)
	if err != nil {
		log.Fatal(err)
	}

	var okj OkJson
	response.GetObject(&okj)

	//var aBool bool
	//_ = json.Unmarshal(okj.Ok[0], &aBool)
	//fmt.Println(aBool)

	var details Details
	_ = json.Unmarshal(okj.Ok[1], &details)
	return details
}

func printResult(response jsonrpc.RPCResponse, err error) {
	if err != nil {
		log.Fatal(err)
	}
	if response.Error != nil {
		log.Println(response.Error.Message)
		return
	}

	j, _ := json.Marshal(response.Result)
	log.Printf("%s\n", j)
}

type CreateAccountPathResult struct {
	Path string `json:"ok"`
}

type AccountsResult struct {
	Accounts []Account `json:"Ok"`
}

type Account struct {
	label string `json:"label"`
	path  string `json:"path"`
}

func RPCAccount(conf GrinConfig) {
	rpcClient := jsonrpc.NewClient(conf.URL)

	//responseCreateAccount, err := rpcClient.Call("create_account_path", "grinclan5")
	//result1 := new(CreateAccountPathResult)
	//err = responseCreateAccount.GetObject(result1)
	//check(err)
	//fmt.Println(result1)
	//printResult(*responseCreateAccount, err)

	responseAccounts, err := rpcClient.Call("accounts")
	result2 := new(AccountsResult)
	err = responseAccounts.GetObject(result2)
	check(err)
	fmt.Printf("%+v\n", result2)
	printResult(*responseAccounts, err)

	//response.GetObject(&resMap)

	//if val, ok := resMap["Ok"]; ok {
	//tup := val.([]interface{})
	//wMap := tup[1].(map[string]interface{})

	//amountAwaitingConfirmation, _ := strconv.ParseInt(wMap["amount_awaiting_confirmation"].(string), 10, 64)
	//amountAwaitingFinalization, _ := strconv.ParseInt(wMap["amount_awaiting_finalization"].(string), 10, 64)
	//amountCurrentSpendable, _ := strconv.ParseInt(wMap["amount_currently_spendable"].(string), 10, 64)
	//amountImmature, _ := strconv.ParseInt(wMap["amount_immature"].(string), 10, 64)
	//amountLocked, _ := strconv.ParseInt(wMap["amount_locked"].(string), 10, 64)
	//lastConfirmedHeight, _ := strconv.ParseInt(wMap["last_confirmed_height"].(string), 10, 64)
	//minimumConfirmations, _ := strconv.ParseInt(wMap["minimum_confirmations"].(string), 10, 64)
	//total, _ := strconv.ParseInt(wMap["total"].(string), 10, 64)

	//walletInfo := WalletInfo{
	//AmountAwaitingConfirmation: amountAwaitingConfirmation,
	//AmountAwaitingFinalization: amountAwaitingFinalization,
	//AmountCurrentSpendable:     amountCurrentSpendable,
	//AmountImmature:             amountImmature,
	//AmountLocked:               amountLocked,
	//LastConfirmedHeight:        lastConfirmedHeight,
	//MinimumConfirmations:       minimumConfirmations,
	//Total:                      total,
	//}

	//jsn, _ := json.Marshal(walletInfo)
	//fmt.Printf("%s\n", jsn)
	//}
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

	summary := RPCSummary(cfg)
	fmt.Println(summary)

	//RPCAccount(cfg)
	//results := AccountsResult{
	//	Accounts: []Account{
	//		{
	//			label: "darkstar",
	//			path:  "1234",
	//		},
	//		{
	//			label: "player",
	//			path:  "12345",
	//		},
	//		{
	//			label: "player3",
	//			path:  "12345",
	//		},
	//	},
	//}
	//b, _ := json.Marshal(results)
	//fmt.Printf("%s", b)
}

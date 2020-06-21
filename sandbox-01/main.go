package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/asciiu/appa/lib/config"
	ecies "github.com/ecies/go"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

func checkErr(label string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", label, err.Error())
		os.Exit(1)
	}
}

type GrinConfig struct {
	URL string `envconfig:"GRIN_API_URL" required:"true"`
}

// type OkJson struct {
// 	Ok []json.RawMessage
// }

// type SummaryInfo struct {
// 	AmountAwaitingConfirmation string `json:"amount_awaiting_confirmation"`
// 	AmountAwaitingFinalization string `json:"amount_awaiting_finalization"`
// 	AmountCurrentlySpendable   string `json:"amount_currently_spendable"`
// 	AmountImmature             string `json:"amount_immature"`
// 	AmountLocked               string `json:"amount_locked"`
// 	LastConfirmedHeight        string `json:"last_confirmed_height"`
// 	MinimumConfirmations       string `json:"minimum_confirmations"`
// 	Total                      string `json:"total"`
// }

func InitSecureApi(conf GrinConfig) ([]byte, error) {
	type Ok struct {
		PublicKey string `json:"Ok"`
	}

	privateKey, err := ecies.GenerateKey()
	if err != nil {
		return []byte{}, fmt.Errorf("generate key failed: %s", err)
	}

	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("init_secure_api", privateKey.PublicKey.Hex(true))
	if err != nil {
		return []byte{}, fmt.Errorf("init_secure_api failed: %s", err)
	}

	var result Ok
	err = response.GetObject(&result)
	if err != nil {
		return []byte{}, fmt.Errorf("get reponse object failed: %s", err)
	}

	log.Infof("received public key result: %s", result.PublicKey)

	remotePublicKey, err := ecies.NewPublicKeyFromHex(result.PublicKey)
	if err != nil {
		return []byte{}, fmt.Errorf("failed remote public key %s", err)
	}
	sharedKey, err := privateKey.ECDH(remotePublicKey)
	if err != nil {
		return []byte{}, fmt.Errorf("failed ecdh %s", err)
	}

	return sharedKey, nil
}

// func GrinSummary(conf GrinConfig) (*SummaryInfo, error) {
// 	rpcClient := jsonrpc.NewClient(conf.URL)
// 	response, err := rpcClient.Call("retrieve_summary_info", true, 10)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var okj OkJson
// 	response.GetObject(&okj)

// 	//var aBool bool
// 	//_ = json.Unmarshal(okj.Ok[0], &aBool)
// 	//fmt.Println(aBool)

// 	var info SummaryInfo
// 	_ = json.Unmarshal(okj.Ok[1], &info)
// 	return &info, nil
// }

func printResult(response jsonrpc.RPCResponse) {
	j, _ := json.Marshal(response)
	log.Printf("%s\n", j)
}

// type ErrAccountExists struct {
// 	Err struct {
// 		AccountLabelAlreadyExists string
// 	}
// }

// type CreateAccountPathResult struct {
// 	Ok string
// }

// type AccountsResult struct {
// 	Ok []struct {
// 		Label string `json:"label"`
// 		Path  string `json:"path"`
// 	}
// }

func EncryptedRquest(conf GrinConfig, nonce, base64Str string) error {
	type params struct {
		Nonce   string `json:"nonce"`
		BodyEnc string `json:"body_enc"`
	}

	p := &params{
		Nonce:   nonce,
		BodyEnc: base64Str,
	}
	j, _ := json.Marshal(p)
	log.Infof("encrypted_request_v3 %s", j)

	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("encrypted_request_v3", p)
	if err != nil {
		return err
	}
	printResult(*response)

	return nil
}

// func GrinAccounts(conf GrinConfig) (*AccountsResult, error) {
// 	rpcClient := jsonrpc.NewClient(conf.URL)
// 	responseAccounts, err := rpcClient.Call("accounts")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var okAccounts AccountsResult
// 	err = responseAccounts.GetObject(&okAccounts)
// 	return &okAccounts, err
// }

// func GrinTransactions(conf GrinConfig) error {
// 	rpcClient := jsonrpc.NewClient(conf.URL)

// 	response, err := rpcClient.Call("retrieve_txs", true, nil, nil)
// 	if err != nil {
// 		return err

// 	}
// 	if response.Error != nil {
// 		return errors.New(response.Error.Message)
// 	}
// 	printResult(*response)

// 	return nil
// }

// func GrinCreateAccount(conf GrinConfig, name string) (string, error) {
// 	rpcClient := jsonrpc.NewClient(conf.URL)

// 	response, err := rpcClient.Call("create_account_path", name)
// 	if err != nil {
// 		return "", err

// 	}
// 	if response.Error != nil {
// 		return "", errors.New(response.Error.Message)
// 	}

// 	var errExists ErrAccountExists
// 	err1 := response.GetObject(&errExists)

// 	var path CreateAccountPathResult
// 	err2 := response.GetObject(&path)

// 	switch {
// 	case err1 != nil:
// 		return "", err1
// 	case err2 != nil:
// 		return "", err2
// 	case errExists.Err.AccountLabelAlreadyExists == name:
// 		return "", errors.New("account already exists with that name")
// 	}

// 	return path.Ok, nil
// }

type Body struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      uint        `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
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
	checkErr("process config", err)

	sharedKey, err := InitSecureApi(cfg)
	checkErr("init secure api", err)

	nonce, nonceHex, err := GenerateNonce()
	checkErr("gen nonce", err)

	body := Body{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "retrieve_summary_info",
		Params:  []interface{}{true, 10},
	}

	req, err := json.Marshal(body)
	log.Infof("new request: %s", req)
	checkErr("marshall json", err)

	base64Str, err := Encrypt(sharedKey, nonce, req)
	checkErr("encrypt message", err)

	err = EncryptedRquest(cfg, nonceHex, base64Str)
	checkErr("what is this?", err)

	//summary, err := GrinSummary(cfg)
	//if err != nil {
	//	log.WithFields(log.Fields{
	//		"function": "GrinSummary",
	//	}).Error(err)
	//} else {
	//	fmt.Println(summary)
	//}

	// path, err := GrinCreateAccount(cfg, "darkstar2")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(path)
	// }

	//accounts, err := GrinAccounts(cfg)
	//if err != nil {
	//	log.WithFields(log.Fields{
	//		"function": "GrinAccount",
	//	}).Error(err)
	//} else {
	//	fmt.Println(accounts)
	//}

	//err = GrinTransactions(cfg)
	//check(err)
}

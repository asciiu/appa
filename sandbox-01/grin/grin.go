package grin

import (
	"encoding/json"
	"fmt"
	"os"

	ecies "github.com/ecies/go"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

func checkErr(label string, err error) {
	if err != nil {
		log.Errorf("%s: %s\n", label, err.Error())
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
		return []byte{}, fmt.Errorf("generate key pair failed: %s", err)
	}

	params := struct {
		Public string `json:"ecdh_pubkey"`
	}{
		Public: privateKey.PublicKey.Hex(true),
	}

	p, _ := json.Marshal(params)
	log.Infof("init_secure_api request %s", p)

	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("init_secure_api", &params)
	if err != nil {
		return []byte{}, fmt.Errorf("init_secure_api failed: %s", err)
	}

	var result Ok
	err = response.GetObject(&result)
	if err != nil {
		return []byte{}, fmt.Errorf("get reponse object failed: %s", err)
	}

	log.Infof("init_secure_api response: %s", result.PublicKey)

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

func EncryptedRquest(conf GrinConfig, nonce []byte, body string) {
	params := struct {
		Nonce   string `json:"nonce"`
		BodyEnc string `json:"body_enc"`
	}{
		Nonce:   fmt.Sprintf("%x", nonce),
		BodyEnc: body,
	}

	j, _ := json.Marshal(params)
	log.Infof("encrypted_request_v3 %s", j)

	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("encrypted_request_v3", &params)
	checkErr("encrypted_request_v3 failed", err)

	printResult(*response)
}

func OpenWallet(conf GrinConfig, key []byte, name, password *string) error {
	body := Body{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "open_wallet",
		Params: struct {
			Name     *string `json:"name"`
			Password string  `json:"password"`
		}{
			Name:     name,
			Password: *password,
		},
	}

	req, err := json.Marshal(body)
	log.Infof("open_wallet request: %s", req)
	checkErr("marshall json", err)

	nonce, err := generateNonce()
	checkErr("failed generate nonce", err)

	base64Str, err := Encrypt(key, nonce, req)
	checkErr("encrypt message", err)

	EncryptedRquest(conf, nonce, base64Str)

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

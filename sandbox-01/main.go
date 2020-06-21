package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/asciiu/appa/lib/config"
	ecies "github.com/ecies/go"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type GrinConfig struct {
	URL string `envconfig:"GRIN_API_URL" required:"true"`
}

type OkJson struct {
	Ok []json.RawMessage
}

type SummaryInfo struct {
	AmountAwaitingConfirmation string `json:"amount_awaiting_confirmation"`
	AmountAwaitingFinalization string `json:"amount_awaiting_finalization"`
	AmountCurrentlySpendable   string `json:"amount_currently_spendable"`
	AmountImmature             string `json:"amount_immature"`
	AmountLocked               string `json:"amount_locked"`
	LastConfirmedHeight        string `json:"last_confirmed_height"`
	MinimumConfirmations       string `json:"minimum_confirmations"`
	Total                      string `json:"total"`
}

func InitSecureApi(conf GrinConfig) (string, error) {
	type Ok struct {
		PublicKey string `json:"Ok"`
	}

	rpcClient := jsonrpc.NewClient(conf.URL)
	secp256k1, err := ecies.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("generate key failed: %s", err)
	}

	response, err := rpcClient.Call("init_secure_api", secp256k1.PublicKey.Hex(true))
	if err != nil {
		return "", fmt.Errorf("init_secure_api failed: %s", err)
	}

	var result Ok
	err = response.GetObject(&result)
	if err != nil {
		return "", fmt.Errorf("get reponse object failed: %s", err)
	}

	remotePublicKey, err := ecies.NewPublicKeyFromHex(result.PublicKey)
	if err != nil {
		return "", fmt.Errorf("failed remote public key %s", err)
	}
	sharedKey, err := secp256k1.ECDH(remotePublicKey)

	pk := ecies.NewPrivateKeyFromBytes(sharedKey)
	ecies.Encrypt(pk.PublicKey)

	return pk.Hex(), nil
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(sharedKey []byte, jsonBody string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(jsonBody)))
	gcm, err := cipher.NewGCMWithNonceSize(block, 12)
	//gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func GrinSummary(conf GrinConfig) (*SummaryInfo, error) {
	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("retrieve_summary_info", true, 10)
	if err != nil {
		return nil, err
	}

	var okj OkJson
	response.GetObject(&okj)

	//var aBool bool
	//_ = json.Unmarshal(okj.Ok[0], &aBool)
	//fmt.Println(aBool)

	var info SummaryInfo
	_ = json.Unmarshal(okj.Ok[1], &info)
	return &info, nil
}

func printResult(response jsonrpc.RPCResponse) {
	j, _ := json.Marshal(response)
	log.Printf("%s\n", j)
}

type ErrAccountExists struct {
	Err struct {
		AccountLabelAlreadyExists string
	}
}

type CreateAccountPathResult struct {
	Ok string
}

type AccountsResult struct {
	Ok []struct {
		Label string `json:"label"`
		Path  string `json:"path"`
	}
}

func GrinAccounts(conf GrinConfig) (*AccountsResult, error) {
	rpcClient := jsonrpc.NewClient(conf.URL)
	responseAccounts, err := rpcClient.Call("accounts")
	if err != nil {
		return nil, err
	}
	var okAccounts AccountsResult
	err = responseAccounts.GetObject(&okAccounts)
	return &okAccounts, err
}

func GrinTransactions(conf GrinConfig) error {
	rpcClient := jsonrpc.NewClient(conf.URL)

	response, err := rpcClient.Call("retrieve_txs", true, nil, nil)
	if err != nil {
		return err

	}
	if response.Error != nil {
		return errors.New(response.Error.Message)
	}
	printResult(*response)

	return nil
}

func GrinCreateAccount(conf GrinConfig, name string) (string, error) {
	rpcClient := jsonrpc.NewClient(conf.URL)

	response, err := rpcClient.Call("create_account_path", name)
	if err != nil {
		return "", err

	}
	if response.Error != nil {
		return "", errors.New(response.Error.Message)
	}

	var errExists ErrAccountExists
	err1 := response.GetObject(&errExists)

	var path CreateAccountPathResult
	err2 := response.GetObject(&path)

	switch {
	case err1 != nil:
		return "", err1
	case err2 != nil:
		return "", err2
	case errExists.Err.AccountLabelAlreadyExists == name:
		return "", errors.New("account already exists with that name")
	}

	return path.Ok, nil
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

	shared, err := InitSecureApi(cfg)
	check(err)

	log.Println(shared)

	summary, err := GrinSummary(cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GrinSummary",
		}).Error(err)
	} else {
		fmt.Println(summary)
	}

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

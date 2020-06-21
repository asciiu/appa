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

func printResult(response jsonrpc.RPCResponse) {
	j, _ := json.Marshal(response)
	log.Printf("%s\n", j)
}

func EncryptedRquest(conf GrinConfig, nonce []byte, base64Str string) error {
	params := struct {
		Nonce   string `json:"nonce"`
		BodyEnc string `json:"body_enc"`
	}{
		Nonce:   fmt.Sprintf("%x", nonce),
		BodyEnc: base64Str,
	}

	j, _ := json.Marshal(params)
	log.Infof("encrypted_request_v3 %s", j)

	rpcClient := jsonrpc.NewClient(conf.URL)
	response, err := rpcClient.Call("encrypted_request_v3", &params)
	if err != nil {
		return err
	}
	printResult(*response)

	return nil
}

func OpenWallet(conf GrinConfig, key []byte, nonce []byte, name, password *string) error {
	body := Body{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "open_wallet",
		Params: struct {
			Name     *string `json:"name"`
			Password *string `json:"password"`
		}{
			Name:     name,
			Password: password,
		},
	}

	req, err := json.Marshal(body)
	log.Infof("open wallet request: %s", req)
	checkErr("marshall json", err)

	base64Str, err := Encrypt(key, nonce, req)
	checkErr("encrypt message", err)

	err = EncryptedRquest(conf, nonce, base64Str)
	checkErr("encrypted request failed", err)

	return nil
}

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

	nonce, err := GenerateNonce()
	checkErr("gen 12 byte nonce", err)

	pass := "I am a warrior"
	OpenWallet(cfg, sharedKey, nonce, nil, &pass)
}

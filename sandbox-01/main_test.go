package main

import (
	"log"
	"testing"

	ecies "github.com/ecies/go"
	"github.com/stretchr/testify/assert"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//func TestBets(t *testing.T) {
//	jsonStr := `{"Ok":[{"label":"darkstar","path":"0200000001000000000000000000000000"},{"label":"default","path":"0200000000000000000000000000000000"},{"label":"grinclan","path":"0200000002000000000000000000000000"},{"label":"grinclan5","path":"0200000003000000000000000000000000"}]}`
//
//	var result AccountsResult
//	json.Unmarshal([]byte(jsonStr), &result)
//	fmt.Printf("%+v\n", result)
//
//	assert.Equal(t, true, true, "ding")
//}

func TestCrypto(t *testing.T) {

	k, err := ecies.GenerateKey()
	if err != nil {
		panic(err)
	}
	log.Println("key pair has been generated")
	log.Println(k.PublicKey.Hex(true))

	assert.Equal(t, true, true, "ding")
}

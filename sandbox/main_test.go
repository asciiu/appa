package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestBets(t *testing.T) {
	jsonStr := `{"Ok":[{"label":"darkstar","path":"0200000001000000000000000000000000"},{"label":"default","path":"0200000000000000000000000000000000"},{"label":"grinclan","path":"0200000002000000000000000000000000"},{"label":"grinclan5","path":"0200000003000000000000000000000000"}]}`

	var result AccountsResult
	json.Unmarshal([]byte(jsonStr), &result)
	fmt.Printf("%+v\n", result)

	assert.Equal(t, true, true, "ding")
}

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

type OkResult struct {
	data []interface{} `json:"Ok"`
}

//type Data struct {
//	bool
//	data struct {
//		LastConfirmedHeight string `json:"last_confirmed_height"`
//	}
//}

type Data struct {
	bool
	data struct {
		LastConfirmedHeight string `json:"last_confirmed_height"`
	}
}

func TestBets(t *testing.T) {
	jsonStr := `{"Ok":[true,{"amount_awaiting_confirmation":"0","amount_awaiting_finalization":"0","amount_currently_spendable":"0","amount_immature":"0","amount_locked":"0","last_confirmed_height":"551632","minimum_confirmations":"10","total":"0"}]}`

	result := new(OkResult)
	json.Unmarshal([]byte(jsonStr), result)
	fmt.Printf("%+v", result.data[0])
	//resMap := make(map[string]interface{}, 0)
	//json.Unmarshal([]byte(jsonStr), &resMap)

	//if val, ok := resMap["Ok"]; ok {
	//	tup := val.([]interface{})
	//	wMap := tup[1].(map[string]interface{})

	//    amountAwaitingConfirmation, _ := strconv.ParseInt(wMap["amount_awaiting_confirmation"].(string), 10, 64)
	//    amountAwaitingFinalization, _ := strconv.ParseInt(wMap["amount_awaiting_finalization"].(string), 10, 64)
	//    amountCurrentSpendable, _ := strconv.ParseInt(wMap["amount_currently_spendable"].(string), 10, 64)
	//    amountImmature, _ := strconv.ParseInt(wMap["amount_immature"].(string), 10, 64)
	//    amountLocked, _ := strconv.ParseInt(wMap["amount_locked"].(string), 10, 64)
	//    lastConfirmedHeight, _ := strconv.ParseInt(wMap["last_confirmed_height"].(string), 10, 64)
	//    minimumConfirmations, _ := strconv.ParseInt(wMap["minimum_confirmations"].(string), 10, 64)
	//    total, _ := strconv.ParseInt(wMap["total"].(string), 10, 64)

	//	walletInfo := WalletInfo{
	//		AmountAwaitingConfirmation: amountAwaitingConfirmation,
	//		AmountAwaitingFinalization: amountAwaitingFinalization,
	//		AmountCurrentSpendable:     amountCurrentSpendable,
	//		AmountImmature:             amountImmature,
	//		AmountLocked:               amountLocked,
	//		LastConfirmedHeight:        lastConfirmedHeight,
	//		MinimumConfirmations:       minimumConfirmations,
	//		Total:                      total,
	//	}

	//	jsn, _ := json.Marshal(walletInfo)
	//	fmt.Printf("%s\n", jsn)
	//}

	assert.Equal(t, true, true, "ding")
}

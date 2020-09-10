package crypto

import (
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
	t.Run("Test WIF", func(t *testing.T) {

		btcWIF, _ := BTC.CreateWIF()
		btcAddress, _ := BTC.GetAddress(btcWIF)
		fmt.Printf("%s - public: %s\n", btcWIF.String(), btcAddress.EncodeAddress())

		ltcWIF, _ := LTC.CreateWIF()
		ltcAddress, _ := LTC.GetAddress(ltcWIF)
		fmt.Printf("%s - public: %s\n", ltcWIF.String(), ltcAddress.EncodeAddress())

		// test that you cannot import a litcoin wif in bitcoin network
		_, err := BTC.ImportWIF(ltcWIF.String())
		assert.Error(t, err, "what is this?")
	})
}

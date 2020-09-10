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

		btcWIF, _ := Bitcoin.CreateWIF()
		btcAddress, _ := Bitcoin.GetAddress(btcWIF)
		fmt.Printf("Key: %s - Pub addr: %s\n", btcWIF.String(), btcAddress.EncodeAddress())

		ltcWIF, _ := Litecoin.CreateWIF()
		ltcAddress, _ := Litecoin.GetAddress(ltcWIF)
		fmt.Printf("Key: %s - Pub addr: %s\n", ltcWIF.String(), ltcAddress.EncodeAddress())

		// test that you cannot import a litcoin wif in bitcoin network
		_, err := Bitcoin.ImportWIF(ltcWIF.String())
		assert.Error(t, err, "what is this?")
	})
}

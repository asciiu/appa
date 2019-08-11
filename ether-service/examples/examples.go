package examples

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func Balance(client *ethclient.Client) {
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	// // reading the latest balance
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}

func BalanceAtBlock(client *ethclient.Client) {
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	blockNumber := big.NewInt(5532993)

	// // reading the latest balance
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}

func GenerateNewWallet() {
	// random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("Private Key: ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public key: ", hexutil.Encode(publicKeyBytes)[4:])

	// Method 1 of obtaining the public address
	address1 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address: ", address1) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// Method 2 of obtaining the public address
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address2 := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println("Address: ", address2)
}

package main

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

// soure: https://www.youtube.com/watch?v=YkJcBvOsMQc
var network = map[string]Network{
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
}

func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParms := &chaincfg.MainNetParams
	networkParms.PubKeyHashAddrID = network.xpubkey
	networkParms.PrivateKeyID = network.xprivatekey
	return networkParms
}

func (network Network) CreateWIF() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}

func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
}

func (network Network) ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	if !wif.IsForNet(network.GetNetworkParams()) {
		return nil, errors.New("The WIF string is not valid for `" + network.name + "` network")
	}
	return wif, nil
}

func main() {
	wif, _ := network["btc"].CreateWIF()
	address, _ := network["btc"].GetAddress(wif)
	fmt.Printf("%s - %s\n", wif.String(), address.EncodeAddress())
	wif, _ = network["ltc"].CreateWIF()
	address, _ = network["ltc"].GetAddress(wif)
	fmt.Printf("%s - %s\n", wif.String(), address.EncodeAddress())

	_, err := network["btc"].ImportWIF(wif.String())
	if err != nil {
		fmt.Println(err.Error())
	}
}

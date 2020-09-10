package crypto

import (
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// Network wraps crypto network info
type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

// Bitcoin network
var Bitcoin = Network{name: "Bitcoin", symbol: "BTC", xpubkey: 0x00, xprivatekey: 0x80}

// Litecoin network
var Litecoin = Network{name: "Litecoin", symbol: "LTC", xpubkey: 0x30, xprivatekey: 0xb0}

// GetNetworkParams returns the chaincfg params
func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParms := &chaincfg.MainNetParams
	networkParms.PubKeyHashAddrID = network.xpubkey
	networkParms.PrivateKeyID = network.xprivatekey
	return networkParms
}

// CreateWIF creates a network specific wallet import format
func (network Network) CreateWIF() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}

// GetAddress returns a pub key address for wif private key
func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
}

// ImportWIF returns a wallet import format from a wifStr
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

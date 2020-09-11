module github.com/asciiu/appa/sandbox-btc

go 1.14

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
)

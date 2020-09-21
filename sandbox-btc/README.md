# btc sandbox

### Running
1. Install [btcd](https://github.com/btcsuite/btcd).
1. `btcd --configfile ./btcd.conf`
1. `go build`
1. `DEBUG=1 ./sandbox-btc


## btcd, btcwallet, and btcctl primer
Starting btcd
```
btcd --configfile ./btcd.conf
```

Starting btcwallet.conf
```
btcwallet -C ./btcwallet.conf
```

Creating an account
```
btcctl -C ./btcctl-wallet.conf createnewaccount alice
```

Unlocking the wallet
```
$ btcctl -C ./btcctl-wallet.conf walletpassphrase PASSPHRASE 3600
```

Generate addresses
```
$ btcctl -C ./btcctl-wallet.conf getnewaddress
MINER_ADDRESS
$ btcctl -C ./btcctl-wallet.conf getnewaddress alice
ALICE_ADDRESS
```

Usinge minging address for btcd
```
$ btcd --configfile ./btcd.conf --miningaddr=MINER_ADDRESS
```

Mining blocks and recieve block rewards
```
btcctl -C ./btcctl.conf generate 100
[...a hundred of hashes...]
$ btcctl -C ./btcctl-wallet.conf getbalance
```

Sending BTC to alice
```
$ btcctl -C ./btcctl-wallet.conf sendtoaddress ALICE_ADDRESS 0.00001
```

## Generated addresses
* miner: SRixeF4rbVRs21EyoD4hcpavbCkPWnjkd4
* alice: SjUnBduscUDnyUpTjmtccXXQZit3YUtnyX
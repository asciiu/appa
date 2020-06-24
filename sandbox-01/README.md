# Resources

* https://github.com/mimblewimble/docs/wiki/Wallet-JSON-RPC-API-Guide
* https://docs.rs/grin_wallet_api/3.1.1/grin_wallet_api/#functions


Running the node and wallet-api:
```
$ grin --floonet
$ grin-wallet --floonet owner_api
```

Do you curl bro?
```
curl -0 -XPOST -u grin:`cat /Users/bishop/Workspace/Rust/grin-wallet/target/release/.owner_api_secret` --data '{"jsonrpc":"2.0","id":1,"method":"encrypted_request_v3","params":{"nonce":"ed09086dee4eb419832f6e75","body_enc":"eYFCvN4hhIYH+5PjH/gBewDZBs+CskIALBFkB76XLG98aJ5oC69F0kBsM2ITFX8tbt6wO4dBQYsfwKiNB5TZEKe/61l4oWokO7y3oZTLetMIrvpQH9nsOdUkXzPJoxqv7MdJanxsGougOJcmb6D3W9kM"}}' http://127.0.0.1:3420/v3/owner

```
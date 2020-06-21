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
curl -0 -XPOST -u grin:`cat /Users/bishop/Workspace/Rust/grin-wallet/target/release/.owner_api_secret` --data '{"jsonrpc":"2.0","method":"encrypted_request_v3","params":{"nonce":"4a8147a78657cdd109f7fda4","body_enc":"SoFHp4ZXzdEJ9/2kwWj7CT3FcEoaBMry2Pr7FxKA1p463SUN65PSwr1trwP3lBwC4K0JlUqQvS80mNvuKjHwCiWzQvbVixWcecvMMsNYSc9ijVD/eS58WICBnEx7+MS3DJ8Pf/co32cNDbChiOzb9Bq1LtQeiLRqIcLsA785"}, "id":1}' http://127.0.0.1:3420/v3/owner

```
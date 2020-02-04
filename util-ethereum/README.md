# How to generate go source from solidity 
1. Generate the application binary interface and evm bytecode. This is needed in 
order to interact with the contract.
```
solc --abi --bin Store.sol -o build
```
2. Use abigen to to generate the go source.
```
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=store.go
```

### Prerequisites
Install ethereum and solidity.

```
brew update
brew tap ethereum/ethereum
brew install ethereum
brew install solidity
```
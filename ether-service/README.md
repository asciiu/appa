# How to generate go source from solidity 
1. Generate the application binary interface and evm bytecode. This is needed in 
order to interact with the contract.
```
solc --abi --bin Contract.sol -o build
```
2. Generate EVN bytecode:
```
solc --bin Solidity.sol > file.bin
```
3. Use abigen to to generate the go source.
```
abigen --bin=./build/Contract.bin --abi=./build/Contract.abi --pkg=store --out=Contract.go
```

### Prerequisites
Install ethereum and solidity.

```
brew update
brew tap ethereum/ethereum
brew install ethereum
brew install solidity
```
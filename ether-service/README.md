# How to generate go source from solidity 
1. Generate the application binary interface. This is needed in 
order to interact with the contract.
```
solc --abi Solidity.sol > file.abi
```
2. Generate EVN bytecode:
```
solc --bin Solidity.sol > file.bin
```
3. Use abigen to to generate the go source.
```
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

### Prerequisites
Install ethereum and solidity.

```
brew update
brew tap ethereum/ethereum
brew install ethereum
brew install solidity
```
# HoldContract

# Ganache Mnemonic

```jsx
electric
gift
gold
bitter
boring
surface
stumble
swift
peanut
adult
horse
dish
```

# Deploy contracts

- Install truffle

```
npm install -g truffle
```

- Move to the smart contracts directory

```jsx
cd
smart - contracts
```

- Deploy the Smart Contracts with truffle

```jsx
truffle
deploy--
reset--
network
development
```

- Generate the necessary files of Geth

```
abigen --abi=./build/contracts/HoldContract.abi --bin=./build/contracts/HoldContract.bin --pkg=holdContract --out=../eth-gateway/holdContract/holdContract.go
```


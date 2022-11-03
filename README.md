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

# Interacting with endpoints

    //users
	//  { id, email, ... }
	//	POST /users (creates a new user)
	// 	GET /users (lists all users)
	// 	GET /users/{id} (obtains the detail of a specific user)
	// 	PUT /users/{id} (replaces a specific user)
	// 	DELETE /users/{id} (removes a specific user)

	//  user wallets
	//	{ id, balance, ..., userId }
	//  POST /users/{userId}/wallets
	//  GET /users/{userId}/wallets
	//  GET /users/{userId}/wallets/{walletId}
	//  ...

	//  games
	//  { id, participants, ... }
	// 	POST /games
	// 	GET /games
	// 	GET /games/{id}
	// 	PUT /games/{id}
	// 	DELETE /games/{id}
	//  POST /games/{id}:join
	//

	//  bets
	//  { id, gameId, userId, amount, ... }
	//  GET /bets?userId&gameId
	//  GET /bets/{id}

```
HTTP req (GET /users) ->
		HTTP infra/framework (== mux) ->
		input infra (== DefaultUsersHttpHandler) ->
			input adapter (== UserInputAdapter) ->
				use case (== user.UseCase) ->
			output adapter (== InMemoryUserStorage) -->
		output infra
```
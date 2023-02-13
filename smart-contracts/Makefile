format:
	npx prettier --write .

clean:
	npx hardhat clean

compile:
	npx hardhat compile

generate:
	npx hardhat run scripts/generate-abi-bin.ts

tests:
	npx hardhat test

tests-gas-report:
	REPORT_GAS=true npx hardhat test

start:
	npx hardhat node

deploy-local:
	npx hardhat run --network localhost scripts/deploy.ts

deploy-docker:
	docker exec -it contablock-blockchain npx hardhat run --network localhost scripts/deploy.ts
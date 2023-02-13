up-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-up.sql

down-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-down.sql

lint:
	golangci-lint run

start-local:
	make gw-client-generator
	go run main.go

start-docker:
	docker run -d -p 4000:4000 --name=contablock-backend contablock-backend

docker-build:
	docker build --tag contablock-backend .

gw-client-generator:
	cd ../smart-contracts && make clean
	cd ../smart-contracts && make compile
	cd ../smart-contracts && make generate
	abigen --bin=../smart-contracts/build/ContaBlockContract.bin --abi=../smart-contracts/build/ContaBlockContract.abi --pkg=contablock --out=pkg/smart-contract-contablock/infra/outputhttp/generated/contablock_gw_client_generated.go
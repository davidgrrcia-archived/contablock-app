up-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-up.sql

down-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-down.sql

clean:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)
	docker system prune

provision:
	docker-compose up -d

lint:
	golangci-lint run

start:
	go run main.go

docker-build:
	docker build --tag contablock-backend .
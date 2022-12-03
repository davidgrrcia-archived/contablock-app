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

start: 
	npm run dev

backend-lint:
	cd backend && golangci-lint run

run-backend:
	cd backend && go run main.go

run-frontend:
	cd frontend && npm run dev

render-api-docs:
	cd docs/open-api && http-server

validate-api-docs:
	cd docs/open-api && spectral lint open-api.yaml
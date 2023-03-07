clean:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)
	docker system prune -f

provision:
	docker-compose up --build -d
	sleep 10
	cd backend && make up-migrations
	cd smart-contracts && make deploy-docker

install:
	cd setup && go run main.go

start-backend:
	cd backend && make start-local

start-frontend:
	cd frontend && make start-local

start-gateway:
	cd gateway && go run main.go

render-api-docs:
	cd docs/openapi && http-server

validate-api-docs:
	cd docs/openapi && spectral lint openapi.yaml
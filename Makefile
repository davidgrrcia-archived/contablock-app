clean:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)
	docker system prune

provision:
	cd backend && make provision
	sleep 10
	cd backend && make up-migrations
	cd smart-contracts && make clean
	cd smart-contracts && make compile
	cd smart-contracts && make generate
	cd backend && make gw-client-generator
	cd smart-contracts && make start

render-api-docs:
	cd docs/openapi && http-server

validate-api-docs:
	cd docs/openapi && spectral lint openapi.yaml

start-backend:
	cd backend && make start

start-frontend:
	cd frontend && make start
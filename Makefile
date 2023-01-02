clean:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)
	docker system prune

provision:
	cd backend && docker-compose up -d

render-api-docs:
	cd docs/openapi && http-server

validate-api-docs:
	cd docs/openapi && spectral lint openapi.yaml

start-backend:
	cd backend && make start

start-frontend:
	cd frontend && make start
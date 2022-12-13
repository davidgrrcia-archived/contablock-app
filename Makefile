clean:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)
	docker system prune

provision:
	docker-compose up -d

render-api-docs:
	cd docs/open-api && http-server

validate-api-docs:
	cd docs/open-api && spectral lint open-api.yaml
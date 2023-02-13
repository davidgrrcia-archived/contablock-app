format:
	npx prettier --write .

lint:
	npm run lint

start-local:
	npm run dev

start-docker:
	make docker-build
	docker run --net=host -d contablock-frontend

docker-build:
	docker build --tag contablock-frontend .
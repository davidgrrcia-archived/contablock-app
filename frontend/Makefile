format:
	npx prettier --write .

lint:
	npm run lint

start-local:
	npm run dev

start-docker:
	docker run -d -p 5173:5173 --name=contablock-frontend contablock-frontend

docker-build:
	docker build --tag contablock-frontend .
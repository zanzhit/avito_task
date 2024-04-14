build:
	docker-compose build banner-app

run:
	docker-compose up banner-app

migrate:
	migrate -path ./migrate -database 'postgres://postgres:12345@0.0.0.0:5432/postgres?sslmode=disable' up
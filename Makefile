up: stop build migrate
	docker-compose up

upp: stop build migrate
	docker-compose up --scale wclient=5

build:
	docker-compose build

dev: build
	docker-compose up -d postgres

stop:
	docker-compose down

test: dev
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

migrate:
	docker-compose up -d postgres
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database test"
	docker-compose exec postgres psql -U postgres -c "create database development"
	docker-compose exec postgres psql -U postgres -c "create database production"
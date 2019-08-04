up: stop build migrate_up
	docker-compose up

upp: stop build migrate_up
	docker-compose up --scale wclient=5

build: deps
	docker-compose build

dev: build
	docker-compose up postgres graphite grafana redis wauth wgame

stop:
	docker-compose down

test: stop migrate
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

migrate_up:
	docker-compose up -d postgres
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database development"

migrate:
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database test"
	docker-compose exec postgres psql -U postgres -c "create database development"
	docker-compose exec postgres psql -U postgres -c "create database production"

deps:
	sh deps.sh
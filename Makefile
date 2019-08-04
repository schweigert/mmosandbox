up: stop build migrate_up
	docker-compose up

upp: stop build migrate_up
	docker-compose up --scale wclient=5

build:
	docker-compose build

dev: build
	docker-compose up -d postgres graphite redis wauth wgame

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

travis: stop_default_services dev migrate

deploy:
	ruby push.rb

stop_default_services:
	# Disable services enabled by default
	# http://docs.travis-ci.com/user/database-setup/#MySQL
	sudo /etc/init.d/mysql stop
	sudo /etc/init.d/postgresql stop
up: stop build migrate_up migrate_redis
	docker-compose up

upp: stop build migrate_up migrate_redis
	docker-compose up --scale wclient=5

build:
	docker-compose build

dev: build migrate
	docker-compose up -d postgres graphite redis
	sleep 10
	docker-compose up -d wauth wgame
	sleep 10

stop:
	docker-compose down

test: stop dev
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

migrate_up:
	docker-compose up -d postgres
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database development"

migrate_redis:
	docker-compose up -d redis
	sleep 5

migrate:
	docker-compose up -d postgres
	sleep 5
	docker-compose exec postgres psql -U postgres -c "create database test"
	docker-compose exec postgres psql -U postgres -c "create database development"
	docker-compose exec postgres psql -U postgres -c "create database production"

deps:
	sh deps.sh

travis_deps:
	docker volume create --name=grafana-volume

travis: travis_deps stop_default_services dev

deploy:
	ruby push.rb

stop_default_services:
	# Disable services enabled by default
	# http://docs.travis-ci.com/user/database-setup/#MySQL
	sudo /etc/init.d/mysql stop
	sudo /etc/init.d/postgresql stop

install_beacon:
	go install github.com/schweigert/mmosandbox/beacon

certs_chmod:
	chmod 400 certs/*

ssh_metrics: certs_chmod
	ssh -i certs/metrics.cert ubuntu@10.20.218.199

ssh_databases: certs_chmod
	ssh -i certs/databases.cert ubuntu@10.20.218.221

ssh_gameservers: certs_chmod
	ssh -i certs/gameservers.cert ubuntu@10.20.218.231
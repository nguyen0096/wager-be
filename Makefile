db-init:
	docker-compose up -d postgres
.PHONY: db-init

db-migrate:
	migrate -path db/migrations -database "postgres://wager:wager@127.0.0.1:5444/wagerdb?sslmode=disable" up
.PHONY: db-migrate

docker-up:
	docker-compose -f docker-compose.yml up -d
.PHONY: docker-up

docker-up-db:
	docker-compose -f docker-compose.yml up -d db
.PHONY: docker-up-db

docker-down:
	docker-compose -f docker-compose.yml down --remove-orphans -v
.PHONY: docker-down

docker-build-local:
	docker build -t wager-be:local . 
.PHONY: docker-build

build:
	go build -o ./bin/wager-be
.PHONY: build

test-unit:
	go test -v -race ./...

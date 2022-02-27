db-init:
	docker-compose up -d postgres
.PHONY: db-init

db-migrate:
	migrate -path db/migrations -database "postgres://wager:wager@127.0.0.1:5444/wagerdb?sslmode=disable" up
.PHONY: db-migrate
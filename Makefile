ifneq (,$(wildcard ./.env))
    include .env
    export
endif

adminer:
	docker run --rm -it --network host adminer

build:
	go build -o app cmd/server/main.go

run:
	docker compose up --build

test:
	go test -v -cover ./...

migrate: 
	docker run --rm -v ./db/migration:/migration --network host --env-file .env migrate/migrate -path /migration/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(SSL_MODE)" --verbose up

migrate-down: 
	docker run --rm -it -v ./db/migration:/migration --network host --env-file .env migrate/migrate -path /migration/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(SSL_MODE)" --verbose down

sqlc:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

postgresql:
	docker run --rm -it --name postgres12 -p $(POSTGRES_PORT):$(POSTGRES_PORT) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:15-alpine


.PHONY: build run test migrate migrate-down sqlc postgresql

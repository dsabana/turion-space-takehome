PG_USER?=turion
PG_PASSWORD?=secret
PG_HOST?=localhost
PG_PORT?=5432
PG_DB?=turion
PG_SCHEMA?=telemetry
DB_CONTAINER?=turion_test_db

.PHONY: pg
pg:
	@echo ++++Checking if a container is already running:
	@docker ps -a -q --filter "name=$(DB_CONTAINER)" | grep -q . && docker rm -f $(DB_CONTAINER) || echo Awesome!! No DB present.
	@echo
	@echo ++++Spinning up a new PostgreSQL container.
	@echo Container Name: $(DB_CONTAINER)
	@docker run -d --name $(DB_CONTAINER) -p $(PG_PORT):5432 -e POSTGRES_USER=$(PG_USER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) postgres:14.1-alpine
	@echo

.PHONY: pg-schema
pg-schema:
	@echo ++++Checking if PostgreSQL container is ready...
	@timeout 90s bash -c "until docker exec $(DB_CONTAINER) pg_isready ; do sleep 1 ; done"
	@sleep 1
	@echo
	@echo ++++Creating Schema
	@docker exec -it $(DB_CONTAINER) psql -U$(PG_USER) -a $(PG_USER) -c 'CREATE SCHEMA IF NOT EXISTS $(PG_SCHEMA);'
	@echo

.PHONY: migrate-up
migrate-up:
	@echo ++++Migrating PostgreSQL Database tables forward
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@migrate -path db/migrations -database "postgresql://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable&search_path=$(PG_SCHEMA)" -verbose up
	@echo

.PHONY: migrate-down
migrate-down:
	@echo ++++Migrating PostgreSQL Database tables backward
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@migrate -path db/migrations -database "postgresql://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable&search_path=$(PG_SCHEMA)" -verbose down
	@echo

.PHONY: local-db
local-db: pg pg-schema migrate-up

.PHONY: run-server
run-server:
	go run ./cmd/telemetryServer

.PHONY: build-server
build-server:
	go build -o telemetryServer ./cmd/telemetryServer

.PHONY: run-ingestion-service
run-ingestion-service:
	go run ./cmd/telemetryIngest

.PHONY: build-ingestion-service
build-ingestion-service:
	go build -o telemetryIngest ./cmd/telemetryIngest

.PHONY: run-telemetry-api
run-telemetry-api:
	go run ./cmd/telemetryApi

.PHONY: build-telemetry-api
build-telemetry-api:
	go build -o telemetryApi ./cmd/telemetryApi

.PHONY: start-frontend
start-frontend:
	@echo ++++ Starting React Frontend
	@npm --prefix ./frontend run dev

.PHONY: dev dev-cli run run-cli build build-api build-cli start cli clean db-up db-down db-logs db-reset fmt lint docs


# Run in development mode with hot reloading
dev: db-up
	air -c air-api.toml

dev-cli: db-up
	air -c air-cli.toml

# Run in development mode (no hot reloading)
run: 
	go run cmd/api/main.go

run-cli:
	go run cmd/cli/main.go

# Build the API and CLI binaries
build: build-api build-cli

build-api: 
	go build -o tmp/api ./cmd/api

build-cli:
	go build -o tmp/cli ./cmd/cli

start: build 
	./tmp/api

cli: build-cli
	./tmp/cli

clean:
	rm -f tmp/api
	rm -f tmp/cli

db-up:
	@docker compose up -d neo4j
	@echo "Waiting for Neo4j to be ready..."
	@until docker exec junction-neo4j cypher-shell -u neo4j -p password123 "RETURN 1" >/dev/null 2>&1; do \
		echo "Neo4j is starting up..."; \
		sleep 2; \
	done
	@echo "Neo4j is ready!"

db-down:
	docker compose down

db-logs:
	docker compose logs -f neo4j

# WARNING: Destroys all neo4j data
db-reset:
	docker compose down -v
	docker compose up -d neo4j

fmt:
	go fmt ./...
	swag fmt ./...

lint:
	golangci-lint run --timeout 5m

docs: 
	swag init -g cmd/api/main.go
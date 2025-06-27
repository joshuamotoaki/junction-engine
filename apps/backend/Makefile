.PHONY: run dev build start clean db-up db-down db-logs db-reset fmt lint

run: db-up
	go run main.go

dev: db-up
	air

build:
	go build -o junction-engine main.go

start: build db-up
	./junction-engine

clean:
	rm -f junction-engine

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

lint:
	golangci-lint run --timeout 5m
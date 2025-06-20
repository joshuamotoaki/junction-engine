run:
	go run main.go

dev:
	air

build:	
	go build -o junction-engine main.go

start: build
	./junction-engine

clean:
	rm -f junction-engine

db-up:
	docker compose up -d

db-down:
	docker compose down

db-logs:
	docker-compose logs -f neo4j

# WARNING: Destroys all neo4j data
db-reset: 
	docker compose down -v
	docker compose up -d neo4j

fmt:
	go fmt ./...

lint:
	golangci-lint run --timeout 5m

.PHONY: run dev build start clean db-up db-down db-logs db-reset fmt lint

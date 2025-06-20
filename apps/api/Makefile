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

.PHONY: run dev build start clean

.PHONY: install run test lint ci

PORT = 8080

install:
	go mod tidy

run: install
	TODO_PORT=$(PORT) go run .

test: install
	go test -v ./...

lint: install
	golangci-lint run

ci: test lint

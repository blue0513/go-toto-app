.PHONY: run test lint ci

PORT = 8080

run:
	TODO_PORT=$(PORT) go run main.go

test:
	go test -v ./...

lint:
	golangci-lint run

ci: test lint

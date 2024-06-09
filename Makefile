APP_NAME=products

.PHONY: build test

build:
	go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o bin/$(APP_NAME) cmd/main.go

test:
	go test ./... -cover -vet all

run:
	go run cmd/main.go
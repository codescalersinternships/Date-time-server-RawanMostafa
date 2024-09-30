build:
	go build -o httpout ./cmd/httpserver/main.go
	go build -o ginout ./cmd/ginserver/main.go

format:
	go fmt ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	golangci-lint run ./... 

build-run-Images:
	docker-compose up

all: build format lint build-run-Images

generate:
	go generate ./...

build:
	go build -o bin/ ./...

run:
	go run main.go
all: clean build test

clean:
	rm -f bin/consolidator
build:
	go build -o bin/consolidator cmd/app/main.go
test:
	go test ./internal/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

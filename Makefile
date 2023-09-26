all: clean build test

clean:
	rm -f bin/consolidator
build:
	go build -o bin/consolidator cmd/app/main.go
test:
	go test ./internal/...
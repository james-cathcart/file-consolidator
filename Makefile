all: clean build

clean:
	rm -f bin/consolidator
build:
	go build -o bin/consolidator cmd/app/main.go

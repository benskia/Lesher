.PHONY: run test build clean

run:
	go run cmd/Lesher/main.go

test:
	go test ./...

build:
	go build -o bin/Lesher cmd/Lesher/main.go

clean:
	rm -rf bin/
	rm -rf $$XDG_CONFIG_HOME/lesher/

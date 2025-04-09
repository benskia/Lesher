.PHONY: run test build clean

run:
	go run cmd/Thresher/main.go

test:
	go test ./...

build:
	go build -o bin/Thresher cmd/Thresher/main.go

clean:
	rm -rf bin/
	rm -rf $$XDG_CONFIG_HOME/lesher/

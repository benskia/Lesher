.PHONY: run test build clean install

run:
	go run ./cmd/Thresher/main.go

build:
	go build -o ./bin/Thresher ./cmd/Thresher/main.go

clean:
	rm -rf ./bin/
	rm -rf $$HOME/.config/Thresher

install:
	go install ./cmd/Thresher/

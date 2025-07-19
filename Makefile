.PHONY: build test clean

build:
	mkdir -p ./bin
	go build -o ./bin/task-manager-cli .

test:
	go test ./...

clean:
	rm -rf ./bin
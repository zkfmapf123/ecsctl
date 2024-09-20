clear:
	rm -rf ./bin/*

build: clear
	go build -o bin/main main.go

lint:
	clear
	golangci-lint run

dev-run: lint
	go run main.go
	
test:
	go test -v ./...

run: build
	./bin/main
clear:
	rm -rf ./bin/*

build: clear
	go build -o bin/ecsctl main.go

lint:
	clear
	golangci-lint run

test:
	go test -v ./...

_build: clear	
	go build -o ecsctl main.go

dev-run: lint 
	make _build
	./main

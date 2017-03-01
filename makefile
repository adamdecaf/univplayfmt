.PHONY: build test vet

build: vet
	go build .

vet:
	go fmt github.com/adamdecaf/univplayfmt
	go tool vet .

test: build
	go test -v ./...

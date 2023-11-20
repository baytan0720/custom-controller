BINARY_NAME=controller

all: go_module build

go_vendor:
	go mod vendor

go_module:
	go mod tidy

build:
	go build -o bin/$(BINARY_NAME) .
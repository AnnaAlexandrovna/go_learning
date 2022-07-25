.PHONY: build
build:
	go build -v ./cmd/app
.PHONY: run
run:
	go run -v ./cmd/app
.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.PHONY: clean
clean:
	go clean -testcache
.DEFAULT_GOAL := build
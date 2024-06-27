build:
	@go build -o bin/blockey

run: build
	@./bin/docker

test:
	@go test -v ./...
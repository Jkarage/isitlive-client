build:
	@go build -o ./bin/client

run: build
	@./bin/client tausi.africa
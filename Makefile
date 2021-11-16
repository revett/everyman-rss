generate:
	@go generate ./...

lint:
	@golangci-lint run

run:
	@go build .
	@./everyman-rss

generate:
	@go generate ./...

run:
	@go build .
	@./everyman-rss

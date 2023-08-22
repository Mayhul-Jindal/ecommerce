build:
	@go build -o bin/book-store-microservice

run: build 
	@./bin/book-store-microservice

test: 
	@go test -v ./...

	
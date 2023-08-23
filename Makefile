DB_URL=postgresql://admin:admin@localhost:5432/book-store?sslmode=disable

postgres:
	docker run -d --rm -p 5432:5432 --name postgres-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin postgres

createdb:
	 docker exec -it postgres-container createdb --username=admin --owner=admin book-store

dropdb:
	docker exec -it postgres-container dropdb -U admin book-store

migrateup:
	migrate -path ./database/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path ./database/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc


# build:
# 	@go build -o bin/book-store-microservice

# run: build 
# 	@./bin/book-store-microservice

# test: 
# 	@go test -v ./...

	
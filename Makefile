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

test: 
	go test -v ./...

run: postgres createdb migrateup
	go run .

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
	
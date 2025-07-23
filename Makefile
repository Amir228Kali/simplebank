migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine3.22

start-postgres:
	docker start postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb --username=root simple_bank

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

server:
	go run main.go

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/Amir228Kali/simplebank/db/sqlc Store

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test mockgen start-postgres migratedown1 migrateup1
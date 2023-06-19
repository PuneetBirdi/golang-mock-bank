postgres:
	docker run --name bank-postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root bank_db

dropdb:
	docker exec -it bank-postgres dropdb --username=root --owner=root bank_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up

awsmigrateup:
	migrate -path db/migration -database "postgresql://root:SDbmjztO2e3uhLIxDeQd@bank-db.cwijsumpbo6j.us-east-1.rds.amazonaws.com:5432/bank_db" -verbose up

migrateupone:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose down

awsmigratedown:
		migrate -path db/migration -database "postgresql://root:SDbmjztO2e3uhLIxDeQd@bank-db.cwijsumpbo6j.us-east-1.rds.amazonaws.com/bank_db" -verbose down

migratedownone:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/PuneetBirdi/golang-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateupone migratedown migratedownone sqlc test server mock

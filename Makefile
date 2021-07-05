def:
	echo "ghel";

migrateup:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5444/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5444/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/lightsOfTruth/dog-walker/db/sqlc Store

.PHONY: migrateup migratedown server sqlc test mock

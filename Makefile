def:
	echo "ghel";

migrateup:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5433/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5433/postgres?sslmode=disable" -verbose down
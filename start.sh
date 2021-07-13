#!/bin/sh

# set -e stops the execution of a script if a command or pipeline has an error
set -e

# DB_SOURCE="postgresql://postgres:pass@localhost:5444/postgres?sslmode=disable"

echo "running db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "starting application"

# run args passed
exec "$@"
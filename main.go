package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lightsOfTruth/dog-walker/api"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:pass@db:5432/postgres?sslmode=disable"
	serverAddress = "0.0.0.0:5000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start serevr", err)
	}

}
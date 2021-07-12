package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lightsOfTruth/dog-walker/api"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
	"github.com/lightsOfTruth/dog-walker/helpers"
)

func main() {
	config, err := helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start serevr", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start serevr", err)
	}

}

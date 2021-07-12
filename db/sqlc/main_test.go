package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/lightsOfTruth/dog-walker/helpers"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, configErr := helpers.LoadConfig("../../")
	if configErr != nil {
		log.Fatal("cannot load config", configErr)
	}

	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

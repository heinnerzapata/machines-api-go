package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://amxolmvh:Ifc6YHUH5PdD1fTrVb0046CpbpuAAJXz@trumpet.db.elephantsql.com/amxolmvh"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	var err error

	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}

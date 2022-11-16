package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

//Tests for local environment only.
const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/simple?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

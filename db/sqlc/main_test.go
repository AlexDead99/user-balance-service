package db

import (
	"database/sql"
	"testing"
	"log"
	"os"
	"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/balance?ssl_mode=false"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot establish connection to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
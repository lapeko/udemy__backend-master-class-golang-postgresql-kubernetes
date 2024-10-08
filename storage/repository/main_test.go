package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbDriverStr = "postgres"
	dbConnStr   = "postgres://root:1234@127.0.0.1:5432/simple_bank_test?sslmode=disable"
	testQueries *Queries
	testDb      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriverStr, dbConnStr)

	if err != nil {
		log.Fatalln(err)
	}

	defer truncateTables()

	testQueries = New(testDb)

	os.Exit(m.Run())
}

func truncateTables() {
	if _, err := testDb.Exec("TRUNCATE users CASCADE"); err != nil {
		log.Fatalln("truncateTables failed", err)
	}
}

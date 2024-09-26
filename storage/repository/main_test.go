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
	db          *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open(dbDriverStr, dbConnStr)

	if err != nil {
		log.Fatalln(err)
	}

	testQueries = New(db)

	defer func() {
		_, err := db.Exec("DELETE FROM accounts")
		if err != nil {
			log.Println("Failed to clean up accounts:", err)
		}
	}()

	exitCode := m.Run()

	os.Exit(exitCode)
}

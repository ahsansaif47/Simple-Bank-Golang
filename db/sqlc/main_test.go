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
	dbSource = "postgresql://postgres:61926114@localhost:5432/simple_bank?sslmode=disable"
)

// defining testQueries a golbally..
// it will be used in unit testing..
var testQueries *Queries

var testDB *sql.DB

// Entry point for all unit tests within package "db"
// Contains unit tests for db CRUD operations..
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Database connection Error!", err)
	}

	// using connection to create testQueries object..
	testQueries = New(testDB)

	os.Exit(m.Run())
}

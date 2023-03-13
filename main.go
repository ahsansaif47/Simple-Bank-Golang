package main

import (
	"database/sql"
	"log"
	"simple_bank_project/api"
	db "simple_bank_project/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:61926114@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Database connection Error!", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("localhost:8081")
	if err != nil {
		log.Fatal("Unable to start server")
	}
}

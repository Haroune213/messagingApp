package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "messagingApp"
)

var db *sql.DB

func OpenDB() (*sql.DB, error) {
	var err error
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		user, password, host, port, dbname, "disable")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

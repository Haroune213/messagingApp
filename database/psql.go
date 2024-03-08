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
	password = "your-password"
	dbname   = "calhounio_demo"
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
	fmt.Println("Connected to the database")

	return db, nil
}

func GetUser(email string) (int, string, string, error) {
	var password string
	var id int
	var username string

	sqlStatement := `SELECT id,username,password FROM users WHERE email=$1`
	row := db.QueryRow(sqlStatement, email)
	switch err := row.Scan(&id, &username, &password); err {
	case sql.ErrNoRows:
		return 0, "", "", err
	case nil:
		return id, username, password, err
	default:
		return id, username, password, err
	}
}

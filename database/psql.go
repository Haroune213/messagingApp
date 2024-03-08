package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	fmt.Println("Connected to the database")

	return db, nil
}

func GetUserValue(email string) (int, string, string, error) {
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

func CreateUserValue(email string, username string, pswd string) (int, bool) {
	var sqlId int

	currentDateTime := time.Now()
	formattedDateTime := currentDateTime.Format("2006-01-02 15:04:05")
	db.QueryRow(`INSERT INTO users (username, password,email,created_at,last_conn)
	VALUES ( $1, $2,$3,$4,$5);`, username, password, email, formattedDateTime, formattedDateTime)

	getId := `select id from users WHERE username=$1 AND email=$2`
	row := db.QueryRow(getId, username, email)
	row.Scan(&sqlId)

	return sqlId, true
}

func DeleteUserValue(email string) {
	db.Exec(`DELETE FROM users WHERE email = $1;`, email)
}

func UpdateLastConnect(id int) {
	currentDateTime := time.Now()
	formattedDateTime := currentDateTime.Format("2006-01-02 15:04:05")

	db.Exec(`UPDATE users SET last_conn = $1 WHERE id = $2;`, formattedDateTime, id)
}

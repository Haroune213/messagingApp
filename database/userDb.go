package database

import (
	"database/sql"
	"time"
)

func GetUserValue(email string) (int, string, error) {
	var password string
	var id int
	var username string

	sqlStatement := `SELECT id,username,password FROM users WHERE email=$1`
	row := db.QueryRow(sqlStatement, email)
	switch err := row.Scan(&id, &username, &password); err {
	case sql.ErrNoRows:
		return 0, "", err
	case nil:
		return id, username, err
	default:
		return id, username, err
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

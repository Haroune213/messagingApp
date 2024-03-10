package database

import (
	"database/sql"
	"messagingApp/structs"
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

func GetUserById(id int) (string, error) {
	var username string

	sqlStatement := `SELECT username FROM users WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&username); err {
	case sql.ErrNoRows:
		return "", err
	case nil:
		return username, err
	default:
		return username, err
	}
}

func GetUserByName(username string) (structs.User, error) {
	var user structs.User
	row := db.QueryRow(`SELECT id FROM users WHERE username=$1`, username)
	switch err := row.Scan(&user.ID); err {
	case sql.ErrNoRows:
		return user, err
	case nil:
		return user, nil
	default:
		return user, nil
	}
}

func GetUsersByName(user string) []structs.User {
	var users []structs.User
	rows, err := db.Query("SELECT id, username FROM users WHERE username ILIKE $1 LIMIT 10;", user+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return users
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

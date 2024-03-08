package models

import (
	"messagingApp/database"
	"time"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	LastConn time.Time
}

func GetUser(email string, pswd string) (User, error) {
	id, usr, pswd, err := database.GetUser(email)

	if err != nil {
		return User{}, err
	}

	user := &User{
		ID:       id,
		Username: usr,
		Password: pswd,
		LastConn: time.Now(),
	}

	return *user, nil
}

func CreateUser() {}

func DeleteUser() {}

func RemoveUser() {}
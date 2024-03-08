package models

import (
	"fmt"
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
	id, usr, pswd, err := database.GetUserValue(email)

	if err != nil {
		fmt.Println(err)
		user := &User{
			ID: 0,
		}
		return *user, err
	}

	user := &User{
		ID:       id,
		Username: usr,
		Password: pswd,
		LastConn: time.Now(),
	}

	return *user, nil
}

func CreateUser(email string, username string, pswd string) (int, bool, error) {
	id, _, _, err := database.GetUserValue(email)
	if id != 0 {

		fmt.Println("this step")
		return 0, false, err
	}

	id, worked := database.CreateUserValue(email, username, pswd)

	if !worked {
		return 0, worked, nil
	}

	return id, worked, nil
}

func DeleteUser() {}

func RemoveUser() {}

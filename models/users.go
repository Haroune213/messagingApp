package models

import (
	"fmt"
	"messagingApp/database"
	"messagingApp/structs"
	"time"
)

func GetUser(email string, pswd string) (structs.User, error) {

	id, usr, err := database.GetUserValue(email)

	if err != nil {
		fmt.Println(err)
		user := &structs.User{
			ID: 0,
		}
		return *user, err
	}

	database.UpdateLastConnect(id)

	user := &structs.User{
		ID:       id,
		Username: usr,
		Email:    email,
		LastConn: time.Now(),
	}

	return *user, nil
}

func CreateUser(email string, username string, pswd string) (int, bool, error) {
	id, _, err := database.GetUserValue(email)
	if id != 0 {
		return 0, false, err
	}

	id, worked := database.CreateUserValue(email, username, pswd)

	if !worked {
		return 0, worked, nil
	}

	return id, worked, nil
}

func DeleteUser(email string) {
	database.DeleteUserValue(email)
}

package models

import (
	"messagingApp/database"
)

func GetChannel(channel_id string, user_id int) User {
	target_user := &User{}

	target := database.GetChannelValue(channel_id, user_id)
	if target == 0 {
		return *target_user
	}
	if target != 0 {
		username, last_conn, err := database.GetUserById(target)
		if err != nil {
			return *target_user
		}

		target_user = &User{Username: username, LastConn: last_conn}
	}
	return *target_user
}

func CreateChannel(user1 int, user2 int) string {
	channel_id := database.CreateChannelValue(user1, user2)

	return channel_id
}

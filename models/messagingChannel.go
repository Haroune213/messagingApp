package models

import (
	"messagingApp/database"
	"messagingApp/structs"
)

func GetChannel(channel_id string, user_id int) (structs.User, bool) {
	target_user := &structs.User{}

	target := database.GetChannelValue(channel_id, user_id)
	if target == 0 {
		return *target_user, false

	}
	if target != 0 {
		username, err := database.GetUserById(target)
		if err != nil {
			return *target_user, false
		}

		target_user = &structs.User{Username: username}
	}
	return *target_user, true
}

func CreateChannel(user1 int, user2 int) string {
	channel_id := database.CreateChannelValue(user1, user2)

	return channel_id
}

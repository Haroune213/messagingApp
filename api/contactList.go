package api

import (
	"messagingApp/database"
	"messagingApp/structs"
)

func SideBarContact(id int) []structs.Message_channel {
	channels := database.GetChannelsList(id)

	for i, channel := range channels {
		channels[i].Target_user, _ = database.GetUserById(channel.Target_id)

	}

	return channels
}

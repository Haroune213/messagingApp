package database

import (
	"database/sql"
	"messagingApp/structs"
)

func GetChannelValue(channel_id string, user_id int) int {
	var usr_1, usr_2, target int

	sqlStatement := `SELECT user1,user2 FROM message_channel WHERE id=$1`
	row := db.QueryRow(sqlStatement, channel_id)
	switch err := row.Scan(&usr_1, &usr_2); err {
	case sql.ErrNoRows:
		target = 0
	case nil:
		if user_id == usr_1 {
			target = usr_2
		}
		if user_id == usr_2 {
			target = usr_1
		}
		if user_id != usr_1 && user_id != usr_2 {
			target = 0
		}
	}

	return target
}

func GetChannelsList(user_id int) []structs.Message_channel {
	var channels []structs.Message_channel
	rows, err := db.Query("SELECT id, user1,user2 FROM message_channel WHERE user1= $1 OR user2= $1 ORDER BY creation_date DESC LIMIT 10;", user_id)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var channel structs.Message_channel
		var usr1, usr2 int
		err := rows.Scan(&channel.ID, &usr1, &usr2)
		if usr1 == user_id {
			channel.Target_id = usr2
		}
		if usr2 == user_id {
			channel.Target_id = usr1
		}

		if err != nil {
			panic(err)
		}
		channels = append(channels, channel)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return channels

}

func CreateChannelValue(user1 int, user2 int) string {
	var id string

	db.QueryRow(`INSERT INTO message_channel (user1, user2)
	VALUES ( $1, $2);`, user1, user2)

	getId := `select id from message_channel WHERE user1=$1 AND user2=$2`
	row := db.QueryRow(getId, user1, user2)
	row.Scan(&id)

	return id
}

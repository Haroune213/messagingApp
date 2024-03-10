package database

import (
	"database/sql"
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

func CreateChannelValue(user1 int, user2 int) string {
	var id string

	db.QueryRow(`INSERT INTO message_channel (user1, user2)
	VALUES ( $1, $2);`, user1, user2)

	getId := `select id from message_channel WHERE user1=$1 AND user2=$2`
	row := db.QueryRow(getId, user1, user2)
	row.Scan(&id)

	return id
}

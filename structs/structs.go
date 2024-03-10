package structs

import "time"

type User struct {
	ID       int
	Username string
	Email    string
	LastConn time.Time
}

type Message_channel struct {
	ID           string
	Target_id    int
	Target_user  string
	Last_message string
	Last_user    int
	Last_date    time.Time
}

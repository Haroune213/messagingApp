package structs

import "time"

type User struct {
	ID       int
	Username string
	Email    string
	LastConn time.Time
}

package models

import "time"

type Session struct {
	Email        string
	LastActivity time.Time
}

type UserNow struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Role      string
}

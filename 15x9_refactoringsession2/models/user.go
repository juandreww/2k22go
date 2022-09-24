package models

import "time"

type Session struct {
	email        string
	lastActivity time.Time
}

type UserNow struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Role      string
}

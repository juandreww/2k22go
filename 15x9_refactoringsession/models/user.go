package models

import "time"

type session struct {
	email        string
	lastActivity time.Time
}

type user struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Role      string
}

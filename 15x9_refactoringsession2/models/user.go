package models

// type session struct {
// 	email        string
// 	lastActivity time.Time
// }

type UserNow struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Role      string
}

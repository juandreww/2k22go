package main

import (
	"net/http"
)


func alreadySignup(req *http.Request) bool {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		return false
	}

	un := dbSessions[cookie.Value]
	_, ok := dbUser[un]
	return ok
}
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

func getUser(req *http.Request) user {
	cookie, err := req.Cookie("session-id")
	var u user
	if err != nil {
		return user{}
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUser[un]
	}

	return u
}
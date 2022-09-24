package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/google/uuid"
)


func alreadySignup(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session-id")
	
	if err != nil {
		return false
	}

	s, ok := dbSessions[cookie.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}
	_, ok = dbUser[s.email]
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	cookie, err := req.Cookie("session-id")
	var u user
	if err != nil {
		uuid := uuid.New()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: uuid.String(),
		}

	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	if s, ok := dbSessions[cookie.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
		u = dbUser[s.email]
	}

	return u
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v)
	}
	fmt.Println("")
}
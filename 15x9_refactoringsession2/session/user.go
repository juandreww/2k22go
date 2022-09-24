package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/15x9_refactoringsession2/models"
)

var dbSessionsCleaned time.Time
var dbUser = map[string]models.UserNow{}

var dbSessions = map[string]models.Session{}

func AlreadySignup(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session-id")

	if err != nil {
		return false
	}

	s, ok := dbSessions[cookie.Value]
	if ok {
		s.LastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}
	_, ok = dbUser[s.Email]
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok
}

func GetUser(w http.ResponseWriter, req *http.Request) models.UserNow {
	cookie, err := req.Cookie("session-id")
	var u models.UserNow
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
		s.LastActivity = time.Now()
		dbSessions[cookie.Value] = s
		u = dbUser[s.Email]
	}

	return u
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

func ShowSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v)
	}
	fmt.Println("")
}

package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/15x9_refactoringsession2/models"
)

var SessionsCleaned time.Time
var DBUser = map[string]models.UserNow{}
var DBSessions = map[string]models.Session{}

const SessionLength int = 30

func AlreadySignup(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session-id")

	if err != nil {
		return false
	}

	s, ok := DBSessions[cookie.Value]
	if ok {
		s.LastActivity = time.Now()
		DBSessions[cookie.Value] = s
	}
	_, ok = DBUser[s.Email]
	cookie.MaxAge = SessionLength
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
	cookie.MaxAge = SessionLength
	http.SetCookie(w, cookie)

	if s, ok := DBSessions[cookie.Value]; ok {
		s.LastActivity = time.Now()
		DBSessions[cookie.Value] = s
		u = DBUser[s.Email]
	}

	return u
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range DBSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(DBSessions, k)
		}
	}
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

func ShowSessions() {
	fmt.Println("********")
	for k, v := range DBSessions {
		fmt.Println(k, v)
	}
	fmt.Println("")
}

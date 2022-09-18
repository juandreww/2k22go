package main

import (
	"github.com/google/uuid"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/session", session)
	http.ListenAndServe(":8080", nil)
}

func session(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if (err != nil) {
		uuid := uuid.New()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
		fmt.Println("added new cookie")
	}
	fmt.Println(cookie)
}


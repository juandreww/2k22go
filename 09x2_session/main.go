package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
	FirstName string
	LastName string
	Email string
}

var dbUser = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/welcome", index)
	mux.HandleFunc("/atthebar", atthebar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to The Lemon Bar. Your used method: ", r.Method)
	cookie, err := r.Cookie("session-id")
	if err != nil {
		uuid := uuid.New()
		cookie = &http.Cookie {
			Name:  "session-id",
			Value: uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func submit(w http.ResponseWriter, r *http.Request) {
	fname :=  r.FormValue("firstname")
	lname := r.FormValue("lastname")
	email := r.FormValue("email")
	uuid := uuid.New()
	dbUser[uuid.String()] = user {
		fname,
		lname,
		email,
	}

	cookie := &http.Cookie {
		Name:  "session-id",
		Value: uuid.String(),
		HttpOnly: true,
		// Secure: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/atthebar", http.StatusSeeOther)
}

func atthebar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		HandleError(w, err)
	}
	u := dbUser[cookie.Value]
	
	if u == (user{}) {
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	}
	
	tpl.ExecuteTemplate(w, "atthebar.gohtml", u)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
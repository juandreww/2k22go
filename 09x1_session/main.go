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
/*----------------------------------------------------------------
	1. template mau bisa panggil semua template - ok
	2. ada input email, firstname,dan last name - ok
	3. bisa submit - ok
	3b. apabila submit, maka register user ke dalam 1 map
	4. apabila submit maka register cookies
	5. apabila refresh maka akan auto redirect ke lemonbar
	6. di lemonbar beri info nama, welcome dst

*/
type user struct {
	FirstName string
	LastName string
	Email string
}

var dbUser = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/submit", submit)
	mux.HandleFunc("/atthebar", atthebar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to The Lemon Bar. Your used method: ", r.Method)
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
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
	Password string
}

/*
	1. user bisa signup
	2. apabila masuk ke signup, isi email, password, firstname, dan lastname
	3. di function signup yang sama, gunakan post dan simpan cookie, serta session
	4. cek apabila sudah signup, maka redirect ke the bar
*/

var dbUser = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	mux := http.DefaultServeMux
	// mux.HandleFunc("/index", index)
	mux.HandleFunc("/atthebar", atthebar)
	mux.HandleFunc("/signup", signup)
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

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUser[un]
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func atthebar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		fmt.Println("here")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		fmt.Println("here2")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	u := dbUser[un]
	tpl.ExecuteTemplate(w, "atthebar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup page: ", r.Method)

	if r.Method == http.MethodPost {
		fname :=  r.FormValue("firstname")
		lname := r.FormValue("lastname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		u := user{fname, lname, email, password,}

		uuid := uuid.New()
		cookie := &http.Cookie {
			Name:  "session-id",
			Value: uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}

		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = fname
		dbUser[fname] = u
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
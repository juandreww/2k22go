package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	Password []byte
}

/*
	
*/

var dbUser = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	mux := http.DefaultServeMux
	// mux.HandleFunc("/index", index)
	mux.HandleFunc("/atthebar", atthebar)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
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
		

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := user{fname, lname, email, bs,}

		uuid := uuid.New()
		cookie := &http.Cookie {
			Name:  "session-id",
			Value: uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}

		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = email
		dbUser[email] = u
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login page: ", r.Method)
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		

		u, ok := dbUser[email]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		fmt.Println(u)
		fmt.Println()
		fmt.Println(ok)

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
	
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"io"
)

var tpl *template.Template
var dbUser = map[string]user{}
var dbSessions = map[string]string{}

type user struct {
	FirstName string
	LastName string
	Email string
	Password []byte
	Role string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("jackywhacky"), bcrypt.MinCost)
	dbUser["jackywhacky@gmail.com"] = user{"jacky","whacky","jackywhacky@gmail.com",bs, "admin"}
}

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)

	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/atthebar", atthebar)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
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

func atthebar(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session-id")

	if (!alreadySignup(r))  {
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
	if u.Role == "007" {
		http.Error(w, "007, please dont come to this bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "atthebar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup page: ", r.Method)

	if r.Method == http.MethodPost {
		fname :=  r.FormValue("firstname")
		lname := r.FormValue("lastname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role")
		

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := user{fname, lname, email, bs, role,}

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
		fmt.Println("password")
		fmt.Println(password)
		fmt.Println()

		u, ok := dbUser[email]
		fmt.Println(u)
		fmt.Println()
		fmt.Println(ok)
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		uuid := uuid.New()
		cookie := &http.Cookie {
			Name:  "session-id",
			Value: uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = email
		http.Redirect(w, r, "/atthebar", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session-id")
	if (!alreadySignup(r))  {
		fmt.Println("here")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	delete(dbSessions, cookie.Value)
	cookie = &http.Cookie{
		Name:   "session-id",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
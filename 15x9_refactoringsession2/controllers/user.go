package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/juandreww/2k22go/15x9_refactoringsession2/models"
	"github.com/juandreww/2k22go/15x9_refactoringsession2/session"
	"golang.org/x/crypto/bcrypt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to The Lemon Bar. Your used method: ", r.Method)
	cookie, err := r.Cookie("session-id")
	if err != nil {
		uuid := uuid.New()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
	}

	var u models.UserNow
	// if un, ok := dbSessions[cookie.Value]; ok {
	// 	u = dbUser[un]
	// }

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func AtTheBar(w http.ResponseWriter, r *http.Request) {
	// cookie, _ := r.Cookie("session-id")
	session.ShowSessions()
	if !session.AlreadySignup(w, r) {
		fmt.Println("here")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	u := session.GetUser(w, r)
	if u.Role == "007" {
		http.Error(w, "007, please dont come to this bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "atthebar.gohtml", u)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup page: ", r.Method)
	session.ShowSessions()
	if r.Method == http.MethodPost {
		fname := r.FormValue("firstname")
		lname := r.FormValue("lastname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role")

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := models.UserNow{fname, lname, email, bs, role}

		uuid := uuid.New()
		cookie := &http.Cookie{
			Name:     "session-id",
			Value:    uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}

		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = models.Session{email, time.Now()}
		dbUser[email] = u
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login page: ", r.Method)
	session.ShowSessions()
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
		cookie := &http.Cookie{
			Name:     "session-id",
			Value:    uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = models.Session{email, time.Now()}
		http.Redirect(w, r, "/atthebar", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session-id")
	session.ShowSessions()
	if !session.AlreadySignup(w, r) {
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

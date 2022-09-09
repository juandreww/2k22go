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
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
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

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
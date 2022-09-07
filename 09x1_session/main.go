package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/submit", submit)
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
}
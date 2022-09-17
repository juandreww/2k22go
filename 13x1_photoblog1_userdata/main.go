package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
	// mux.HandleFunc("/read", read)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user-data")
	if (err != nil) {
		uuid := uuid.New()
		cookie = &http.Cookie{
			Name:	"user-data",
			Value:	uuid.String(),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, cookie)
		fmt.Println("added new cookie")
	}
	tpl.ExecuteTemplate(w, "index.gohtml", cookie)
}
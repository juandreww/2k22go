package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
	"strings"
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
	c := getCookie(w, r)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)

	// tpl.ExecuteTemplate(w, "index.gohtml", cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("user-data")
	if err != nil {
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
	
	return cookie
}
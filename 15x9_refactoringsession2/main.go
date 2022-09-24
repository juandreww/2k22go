package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/juandreww/2k22go/15x9_refactoringsession2/models"
)

var tpl *template.Template
var dbSessionsCleaned time.Time
var dbUser = map[string]models.UserNow{}

// var dbSessions = map[string]models.session{}

func main() {
	// mux := http.DefaultServeMux
	// mux.HandleFunc("/index", index)
	// mux.HandleFunc("/atthebar", atthebar)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

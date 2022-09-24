package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/juandreww/2k22go/15x9_refactoringsession2/controllers"
	"github.com/juandreww/2k22go/15x9_refactoringsession2/models"
)

var tpl *template.Template
var dbSessionsCleaned time.Time
var dbUser = map[string]models.UserNow{}

var dbSessions = map[string]models.Session{}

func main() {
	ctl := controllers.NewUserController(tpl)
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", ctl.Index)
	mux.HandleFunc("/atthebar", ctl.AtTheBar)
	mux.HandleFunc("/signup", ctl.SignUp)
	mux.HandleFunc("/login", ctl.Login)
	mux.HandleFunc("/logout", ctl.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

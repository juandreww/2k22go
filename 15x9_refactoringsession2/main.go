package main

import (
	"html/template"
	"net/http"

	"github.com/juandreww/2k22go/15x9_refactoringsession2/controllers"
)

var tpl *template.Template

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

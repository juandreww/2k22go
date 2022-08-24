package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("go.golangscript"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "go.golangscript", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
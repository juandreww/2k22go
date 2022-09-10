package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("hello")
}
package main

import (
	"fmt"
	"os"
	"text/template"
)

type course struct {
	Number string
	Name string
	Units string
}

type semester struct {
	Term string
	Course []course
}

type year struct {
	AcaYear string
	Fall semester
	Spring semester
	Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	
}
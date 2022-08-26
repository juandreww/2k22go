package main

import (
	"log"
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
	Courses []course
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
	years := []year {
		year {
			AcaYear: "2025-2026",
			Fall: semester {
				Term: "Fall",
				Courses: []course {
					course{"CSCI-40", "Introduction to Programming in GO", "4"},
					course{"CSCI-130", "Introduction to Web Programming with GO", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester {
				Term: "Spring",
				Courses: []course {
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-130", "Advanced Web Programming with Go", "5"},
					course{"CSCI-140", "Advanced Mobile Apps with Go", "5"},
				},
			},
		},
		year {
			AcaYear: "2026-2027",
			Summer: semester {
				Term: "Summer",
				Courses: []course {
					course{"CSCI-40", "Introduction to Programming in GO", "4"},
					course{"CSCI-130", "Introduction to Web Programming with GO", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Fall: semester {
				Term: "Fall",
				Courses: []course {
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-130", "Advanced Web Programming with Go", "5"},
					course{"CSCI-140", "Advanced Mobile Apps with Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
package main

import (
	"log"
	"os"
	"text/template"
	"strings"
	// "reflect"
)

type Commodity struct {
	Location, Name string
}

func createCommodityList(data [][]string) []Commodity {
	var list []Commodity
	for k, v := range data {
		if (k > 0) {
			var com Commodity
			for k2, v2 := range v {
				if k2 == 0 {
					com.Location = v2
				} else if k2 == 1 {
					com.Name = v2
				}
			}
			list = append(list, com)
		}
	}
	return list
}

func (p Commodity) marketLocation() string {
	if strings.Contains(p.Name, "Daging") {
		return "ada di Pasar Kopo"
	} else if strings.Contains(p.Name, "Beras") {
		return "ada di Pasar Pademangan"
	}
	return "none"
}

func (p Commodity) marketNumber() int {
	return 7
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := Commodity {
		"Bandung Barat",
		"Cilegon",
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
package main

import (
	"fmt"
	"encoding/csv"
	"text/template"
	"os"
	"log"
	// "reflect"
)

type Commodity struct {
	Location, Name string
}

var tpl *template.Template

// var fm = template.FuncMap{
// 	"add": func(a, b int) int {
// 		return a + b
// 	},
// }

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("hargakomoditas.csv")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", f)
	if err != nil {
		log.Fatalln(err)
	}
	
	
}
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"
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

	commodityList := createCommodityList(data)

	for _, v := range commodityList {
		fmt.Printf("%s: %s\n", v.Location, v.Name)
	}
	
	fmt.Println()

	sort.SliceStable(commodityList, func (i, j int) bool {
		return commodityList[i].Name < commodityList[j].Name
	})

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", commodityList)
	if err != nil {
		log.Fatalln(err)
	}
}
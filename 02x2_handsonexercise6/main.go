package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"encoding/csv"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type Commodity struct {
	Location, Name string
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
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

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
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

	err = tpl.Execute(os.Stdout, commodityList)
	if err != nil {
		log.Fatalln(err)
	}
}
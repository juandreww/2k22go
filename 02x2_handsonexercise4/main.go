package main

import (
	"fmt"
	"text/template"
	"os"
	"log"
	// "reflect"
)
type Item struct {
	Name string
	Quantity int
}

type MealPeriod struct {
	Name string
	Item []Item
}

type DQCoffee struct {
	MealPeriod []MealPeriod
}

type Branch struct {
	BranchName string
	DQCoffee []DQCoffee
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// DQ Coffee 24 / 7
	dqc := Branch {
		"Batam Centre",
		[]DQCoffee {
			DQCoffee {
				[]MealPeriod {
					MealPeriod{
						"Breakfast",
						[]Item {
							Item {
								"Telur Rebus",
								10,
							},
							Item {
								"Ayam Rebus",
								10,
							},
							Item {
								"Sosis Goreng",
								10,
							},
						},
					},
					MealPeriod{
						"Lunch",
						[]Item {
							Item {
								"Ayam Penyet",
								10,
							},
							Item {
								"Tahu, Tempe, Telur",
								10,
							},
							Item {
								"Indomie Goreng",
								10,
							},
							Item {
								"Indomie Kuah",
								10,
							},
						},
					},
					MealPeriod{
						"Dinner",
						[]Item {
							Item {
								"Teh Susu",
								10,
							},
							Item {
								"Kopi Susu",
								10,
							},
							Item {
								"Matcha Latte",
								10,
							},
							Item {
								"Wedang Jahe",
								10,
							},
						},
					},
				},
			},
			DQCoffee {
				[]MealPeriod {
					MealPeriod{
						"Breakfast",
						[]Item {
							Item {
								"Nasi Kucing",
								10,
							},
							Item {
								"Nasi Lemak",
								10,
							},
							Item {
								"Bubur Ayam",
								10,
							},
						},
					},
					MealPeriod{
						"Lunch",
						[]Item {
							Item {
								"Soto Solo",
								10,
							},
							Item {
								"Soto Daging",
								10,
							},
							Item {
								"Soto Rempah",
								10,
							},
							Item {
								"Soto Kecap",
								10,
							},
						},
					},
					MealPeriod{
						"Dinner",
						[]Item {
							Item {
								"Teh Milkshake",
								10,
							},
							Item {
								"Coffee Milk",
								10,
							},
							Item {
								"Coffee Latte",
								10,
							},
							Item {
								"Coffee Arabica",
								10,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(dqc)

	err := tpl.Execute(os.Stdout, dqc)
	if err != nil {
		log.Fatalln(err)
	}
	
	
}
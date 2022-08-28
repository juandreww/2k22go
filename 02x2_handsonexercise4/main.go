package main

import (
	"fmt"
	// "text/template"
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

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
// }

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
		},
	}

	fmt.Println(dqc)

	// nasgor := 5;
	// for k, v := range dq.MealPeriod {
	// 	fmt.Println("MealPeriod: " + v.Name)
	// 	for k2, v2 := range v.Item {
	// 		fmt.Printf("Item: %s with Stocks: %d\n", v2.Name, v2.Quantity)
	// 		if v2.Name == "Indomie Kuah" {
	// 			dq.MealPeriod[k].Item[k2].Quantity -= nasgor
	// 			fmt.Printf("Sisa %s adalah %d\n", v2.Name, v2.Quantity) // disini ga berubah
	// 		}
	// 	}
	// }

	// err := tpl.Execute(os.Stdout, )
	
	
}
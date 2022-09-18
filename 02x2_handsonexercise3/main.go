package main

import (
	"fmt"
	// "reflect"
)

type Item struct {
	Name string
	Quantity int
}

type Period struct {
	Name string
	Item []Item
}

type DQCoffee struct {
	Period []Period
}

func main() {
	// DQ Coffee 24 / 7
	dq := DQCoffee {
		[]Period {
			Period{
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
			Period{
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
			Period{
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
	}

	nasgor := 5;

	// fmt.Println(dq.Period)
	for k, v := range dq.Period {
		fmt.Println("Period: " + v.Name)
		for k2, v2 := range v.Item {
			fmt.Printf("Item: %s with Stocks: %d\n", v2.Name, v2.Quantity)
			if v2.Name == "Indomie Kuah" {
				dq.Period[k].Item[k2].Quantity -= nasgor
				fmt.Printf("Sisa %s adalah %d\n", v2.Name, v2.Quantity) // disini ga berubah
			}
		}
	}

	fmt.Println(dq.Period)
	
	
}
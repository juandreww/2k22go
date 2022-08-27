package main

import (
	"fmt"
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

	fmt.Println(dq)
}
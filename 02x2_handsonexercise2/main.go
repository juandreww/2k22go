package main

import (
	"fmt"
	// "log"
	// "os"
)

type Hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

type Region struct {
	Name string
	Hotel []Hotel
}

type California struct {
	Region []Region
}

func main() {
	calf := California {
		Region: []Region {
			Region {
				"Southern", 
					[]Hotel {
						Hotel {
							"Hotel 81",
							"California A11",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Park",
							"California A12",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Grand",
							"California A13",
							"California South",
							"440004",
							"Southern",
						},
					},
				},
			Region {
				"Southern", 
					[]Hotel {
						Hotel {
							"Hotel 81",
							"California A11",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Park",
							"California A12",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Grand",
							"California A13",
							"California South",
							"440004",
							"Southern",
						},
					},
				},
			Region {
				"Southern", 
					[]Hotel {
						Hotel {
							"Hotel 81",
							"California A11",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Park",
							"California A12",
							"California South",
							"440004",
							"Southern",
						},
						Hotel {
							"Hotel Grand",
							"California A13",
							"California South",
							"440004",
							"Southern",
						},
					},
				},
		},
	}


	fmt.Println(calf)
}
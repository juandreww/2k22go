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
				"Central", 
					[]Hotel {
						Hotel {
							"Hotel Moyes",
							"California B88",
							"California Central",
							"330999",
							"Central",
						},
						Hotel {
							"Hotel January",
							"California C90",
							"California Central",
							"330999",
							"Central",
						},
						Hotel {
							"Hotel Grand",
							"California C91",
							"California Central",
							"330999",
							"Central",
						},
					},
				},
			Region {
				"Northern", 
					[]Hotel {
						Hotel {
							"Hotel Velocity",
							"California VV11",
							"California North",
							"993000",
							"Northern",
						},
						Hotel {
							"Hotel Vienna",
							"California VV13",
							"California North",
							"993000",
							"Northern",
						},
						Hotel {
							"Hotel California",
							"California VV14",
							"California North",
							"993000",
							"Northern",
						},
					},
				},
		},
	}


	fmt.Println(calf)
}
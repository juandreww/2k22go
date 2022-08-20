package main

import (
	"fmt"
)

/*
	tentukan nama pacar
	hitung ukuran kotak kado untuk pacar
	hitung ukuran cincin untuk pacar
	print out diameter jari manis pacar
	print out keliling cincin, dan luas area kotak
*/

type girlfriend struct {
	fName string
	lName string
}

type rings struct {
	radius float64
}

func (r rings) diameter() {
	fmt.Printf("%.2f cm\n", r.radius * 2)
}

func (r rings) circumference() float64 {
	
}

type size interface {
	circumference()
}

func calculateRing(r size) {
	
} 

func main() {
	myGf := girlfriend{
		"Yomelia",
		"Christine",
	}
	
	fmt.Println("My girlfriend name is: ")
	fmt.Println(myGf)
	fmt.Println("I am going to propose her with a ring..")

	r := rings{4}
	fmt.Println("Her finger diameter is ")
	r.diameter()
	fmt.Println("Her finger circumference is ")


	// create type square

	
	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}
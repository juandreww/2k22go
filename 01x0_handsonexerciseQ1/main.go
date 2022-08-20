package main

import (
	"fmt"
	"math"
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
	return math.Round(math.Pi * (r.radius * 2) * 100) / 100
}

type size interface {
	circumference() float64
}

func calculateRing(r size) {
	fmt.Println("Her finger circumference is", r.circumference())
} 

type box struct {
	size float64
}

func (b box) volume() {
	fmt.Printf("%.2f cm", math.Pow(b.size, 3))
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
	calculateRing(r)

	fmt.Println("I have a box, with volume ")
	box := box{10}
	box.volume()

	// create type square

	
	// http.HandleFunc("/", foo)
	// http.HandleFunc("/dog", bar)

	// fmt.Println("Serving...");
	// http.ListenAndServe(":8080", nil)
}
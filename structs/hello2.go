package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

// func main() {
// 	fmt.Println("Hello 2")
// 	Anything(2.55)
// 	Anything("Jamal")
// 	Anything(struct{} {})

// 	myMap := make(map[string]interface{})
// 	myMap["name"] = "Jamal Malu"
// 	myMap["age"] = 15
// 	fmt.Println(myMap)
// }
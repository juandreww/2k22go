package main

import "fmt"

func main() {
	fmt.Println("Hello UK")

	f := true
	flag := &f

	fmt.Println(*flag, flag, &flag)
	if *flag {
		fmt.Println("True")
	} else if !*flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []string{"Yangon", "Manila", "Haiti"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Jamal"
	mymap["age"] = 25

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}
}
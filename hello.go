package main

import (
	"fmt"
)

func main() {
	task()
}

func task() {
	arr := []int{1,2,3,4}
	arr2 := []string{"mang", "tas", "net", "work"}
	arr = append(arr, 1,2,3)
	arr2 = append(arr2, "1","2","3", "jasper")
	fmt.Println(arr, arr2)
}
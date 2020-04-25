package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("My Slice", mySlice)

	for index, i := range mySlice {
		fmt.Println("For Range Of Index", index, "::Value of the Index >>", i)
	}
	newSlice := []int{10, 20, 30}
	together := append(mySlice, newSlice...)
	fmt.Println(together)

}

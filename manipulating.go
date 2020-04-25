package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	// Update the value of map
	testMap["A"] = 100
	testMap["F"] = 1989
	fmt.Println(testMap)
}

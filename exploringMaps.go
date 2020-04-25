package main

import "fmt"

func main() {
	value := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for key, value := range value {
		fmt.Println("Value of Map", key, value)
	}
}

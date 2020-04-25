package main

import "fmt"

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}
	dockerDeepDive := courseMeta{
		Author: "Bishwajit",
		Level:  "Intermediate",
		Rating: 2.8,
	}
	fmt.Println("\n Course Author is::", dockerDeepDive.Author)
	fmt.Println("\n Rating of the Course is::", dockerDeepDive.Rating)
}

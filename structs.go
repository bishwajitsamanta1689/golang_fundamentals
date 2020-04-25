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
		Rating: 5,
	}
	fmt.Println(dockerDeepDive)
}

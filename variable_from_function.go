package main

import (
	"fmt"
	"reflect"
)

func changeCourse(course string) string {
	course = "Go Fundamentals"
	fmt.Println("Trying to Change course", course)
	return course
}
func main() {
	name := "bishwajit"
	course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("\nHi", name, "You are currently watching ", course, "of type,", reflect.TypeOf(name))
	fmt.Println("\nHi", "You are looking at Module", *ptr, "of type,", reflect.TypeOf(*ptr))

	changeCourse(course)
	fmt.Println("\nHi", name, "You are now watching", course)
}

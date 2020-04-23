package main

import (
	"fmt"
	"reflect"
)

// Passing By Pointer Reference
// course is a pointer to the string
func modifyCourse(course *string) string {
	*course = "Go Fundamentals"                           // If we dont use * here, it will say cannot use type string as *string
	fmt.Println("Trying to Change the Course::", *course) // if we dont use * here, it will print the memory address
	return *course
}
func main() {
	name := "bishwajit"
	course := "Docker Deep Dive"
	module := 3.4

	fmt.Println("\nHi", name, "you are currently watching ", course, "of type", reflect.TypeOf(course))
	fmt.Println("\nHi", name, "you are currently watching module", module, "of type", reflect.TypeOf(module))
	modifyCourse(&course)
}

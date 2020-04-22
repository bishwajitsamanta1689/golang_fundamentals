package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Nigel"
	course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("Name is ", name, "is of type::", reflect.TypeOf(name))
	fmt.Println("Course is ", course, "is of type::", reflect.TypeOf(course))
	fmt.Println("Module is ", module, "is of type::", reflect.TypeOf(module))
	fmt.Println("Memory address of ptr ", ptr, "have value of", *ptr)
}

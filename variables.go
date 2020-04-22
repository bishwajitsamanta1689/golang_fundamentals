package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "apple"
	course = "docker deep dive"
	module = 2.3
)

func main() {
	fmt.Println("Name is ", name, "and is of type ::", reflect.TypeOf(name))
	fmt.Println("Module is ", module, "and is of type ::", reflect.TypeOf(module))
	fmt.Println("course is ", course, "and is of type ::", reflect.TypeOf(course))
}

package main

import (
	"fmt"
	"reflect"
)

var (
	a = 10.999
	b = 3
)

func main() {
	fmt.Println("A is ", a, "is of type ::", reflect.TypeOf(a))
	fmt.Println("B is ", b, "is of type ::", reflect.TypeOf(b))
	calc := int(a) + b
	fmt.Println("Addition of the two variables :: ", calc, ".Calc is of type ::", reflect.TypeOf(calc))
}

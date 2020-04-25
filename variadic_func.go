package main

import "fmt"

// Variadic Functions
// A Variable having a Function with n numbers of arguments.
// Create a Function with ellipses which can take n number of arguments.
// Assign the first instance of the argument to a variable
// Find across the range of the Variable and assign the lowest value and return it back.

func main() {
	bestFinish := bestFinishLeague(2, 15, 12, 13, 16, 18, 5, 1)
	fmt.Println(bestFinish)
}

// ... is known as ellipses in Golang, it means it can take any number of inputs
func bestFinishLeague(Finishes ...int) int {
	best := Finishes[0]
	for _, i := range Finishes {
		if i < best {
			best = i
		}
	}
	return best
}

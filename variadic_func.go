package main

import "fmt"

//Take any number of inputs and Returns the lowest

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

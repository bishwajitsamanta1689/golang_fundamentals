package main

import "fmt"

// Using Multiple For Loops one inside another
// Assign 1 variable with a list of items
// Assign another variable with another list of items
// Match if any of the items are similar
func main() {
	courseInProgress := []string{"Docker Deep Dive", "Go Fundamentals", "Docker and Kubernetes"}
	courseCompleted := []string{"Go Fundamentals", "Docker Basic", "C++ Beginners"}

	// Break Labels are used to exit multiple for loops break and go to the top of the code.

checkPoint:
	for _, i := range courseInProgress {
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("Similar Course Found::", j)
				break checkPoint
			}
		}
	}
}

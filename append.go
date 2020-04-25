package main

import "fmt"

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Capacity of the Array %d", cap(mySlice))

	for i := 0; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity of New Array %d", cap(mySlice))
	}

}

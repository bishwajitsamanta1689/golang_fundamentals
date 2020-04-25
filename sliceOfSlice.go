package main

import "fmt"

func main() {
	mySlice := []int{12, 13, 14, 15, 16, 17}
	fmt.Println(mySlice[4])
	sliceOfSlice := mySlice[2:5]
	sliceTillLast := sliceOfSlice[:2]
	fmt.Println(sliceOfSlice)
	fmt.Println(sliceTillLast)
}

package main

import "fmt"

func main() {
	if firstRank, secondRank := 48, 50; firstRank > secondRank {
		fmt.Println("First Rank is greater than Second Rank")
	} else {
		fmt.Println("Second Rank is greater thank First Rank")
	}
}

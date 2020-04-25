package main

import (
	"fmt"
	"time"
)

//Countdown Timer for the ship getting self destruct
func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}

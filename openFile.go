package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("Go_Fundamentals/test.go")
	if err != nil {
		log.Fatalln("Error is", err)
	} else {
		fmt.Println("File Opened successfully !!")
	}

}

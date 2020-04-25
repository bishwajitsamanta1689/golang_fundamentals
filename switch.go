package main

import "fmt"

func main() {
	switch "Docker" {
	case "Linux":
		fmt.Println("\n Here are some Linux Courses!!")
	case "Docker":
		fmt.Println("\n Here are some Docker Courses")
	case "microservices":
		fmt.Println("\n Here are some microservices Courses")
	default:
		fmt.Println("You need to Find among the latest 100 Courses ")
	}
}

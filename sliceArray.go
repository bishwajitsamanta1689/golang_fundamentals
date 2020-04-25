package main

import "fmt"

func main() {
	courseCompleted := []string{"Docker", "Puppet", "Ansible"}
	fmt.Printf("Length of the course %d.\nCourse Capacity is %d", len(courseCompleted), cap(courseCompleted))
}

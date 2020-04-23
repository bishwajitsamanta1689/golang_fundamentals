package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println("\nHi", os.Getenv("USERNAME"), "how's your day going in your laptop", os.Getenv("COMPUTERNAME"), ".I found you are using", runtime.GOOS, "system.")
}

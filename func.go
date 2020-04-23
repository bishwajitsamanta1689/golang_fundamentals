package main

import (
	"fmt"
	"strings"
)

//Make the module titleCase()
// Make the Author titleUpper()
func main() {
	Module := "function basics"
	Author := "bishwajeett Uday samanta"
	fmt.Println(modify(Module, Author))
}
func modify(Module, Author string) (s1, s2 string) {
	Module = strings.Title(Module)
	Author = strings.ToUpper(Author)
	return Module, Author
}

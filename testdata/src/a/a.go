package a

import (
	"fmt"
	"log"
)

func f() {
	// The pattern can be written in regular expression.
	var n1 = 1
	var n2 = 3
	result := add(n1, n2)
	fmt.Println(result) // want "fmt package is used!"
	log.Println(result) // fmt package is not used .

	fmt.Printf("%v", result)                 // want "fmt package is used!"
	fmt.Print(result)                        // want "fmt package is used!"
	callback(func() { fmt.Println(result) }) // want "fmt package is used!"
}

func add(n1, n2 int) int {
	return n1 + n2
}

func callback(f func()) {
	f()
}

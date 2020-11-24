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
	fmt.Println(result) // want "use!"
	log.Println(result) // this is not used fmt package.

	fmt.Printf("%v", result)                 // want "use!"
	fmt.Print(result)                        // want "use!"
	callback(func() { fmt.Println(result) }) // want "use!"
}

func add(n1, n2 int) int {
	return n1 + n2
}

func callback(f func()) {
	f()
}

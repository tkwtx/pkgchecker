package a

import (
	foo "fmt"
	"log"
	_ "reflect"
)

func f() {
	// The pattern can be written in regular expression.
	var n1 = 1
	var n2 = 3
	result := add(n1, n2)
	foo.Println(result) // want "use foo.Println"
	log.Println(result) // This don't output "log" because flag is set "foo".

	foo.Printf("%v", result)                 // want "use foo.Printf"
	foo.Print(result)                        // want "use foo.Print"
	callback(func() { foo.Println(result) }) // want "use foo.Println"
}

func add(n1, n2 int) int {
	return n1 + n2
}

func callback(f func()) {
	f()
}

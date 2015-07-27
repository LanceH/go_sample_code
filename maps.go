package main

import (
	"fmt"
)

// The type of m below is map[string]string
// the key and value types are both part of the type also
var m = make(map[string]string)

func main() {
	fmt.Println(m) // nil map
	m["a"] = "A"
	fmt.Println(m)
	fmt.Println(m["a"])
}

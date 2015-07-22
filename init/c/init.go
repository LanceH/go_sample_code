package c

import (
	"fmt"
)

func init() {
	fmt.Println("init() in package c")
}

// X in c is the version of X() that finally does something
func X(x string) {
	fmt.Println(x)
}

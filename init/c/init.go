package c

import (
	"fmt"
)

func init() {
	fmt.Println("init() in package c")
}

func X(x string) {
	fmt.Println(x)
}

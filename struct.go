package main

/*
Create and use a Struct.
*/

import (
	"fmt"
)

// Car is an example of a struct.
// When it is capitalized it is exported and available outside the package.
type Car struct {
	Speed int
}

// car is not exported but is still available in the same package.
type car struct {
	speed int
}

func main() {
	c1 := Car{44}
	fmt.Println(c1.Speed)
	c2 := Car{Speed: 55}
	fmt.Println(c2.Speed)

	c3 := car{speed: 22}
	fmt.Println(c3.speed)
}

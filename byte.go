package main

import "fmt"

func main() {
	var a byte = 3
	var b uint8 = 4

	// byte is an alias for uint8
	a = b
	fmt.Println(a)
}

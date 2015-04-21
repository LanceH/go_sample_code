package main

/*
init() is a special function called before main.

I managed to read two books before learning this one.
*/


/* 
Interestingly init in b runs before a
in spite of a being listed first.
Switch them in the import list and a will run first.
*/

/*
a and b both require c
c only calls init() once
*/ 

import (
	"./a"
	"./b"
	"fmt"
)

func main() {
	fmt.Println("Main")
	a.X()  // have to call something so a and b
	b.X()  // are used and don't throw an error
}

func init() {
	fmt.Println("Init Main")
}

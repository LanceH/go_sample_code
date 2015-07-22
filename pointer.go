package main

import (
	"fmt"
)

func main() {
	a := 77
	b := &a
	fmt.Println(a)  // a is a normal int value (77)
	fmt.Println(b)  // b is a pointer to a (0x20818a220 or something like that)
	fmt.Println(*b) // 77

	c := C{z: 77}
	fmt.Println(c)   // c is a struct
	fmt.Println(c.z) // c.z is the value 77 again
	d := &c
	fmt.Println(d)      // d is a pointer to a struct
	fmt.Println(*d)     // *d is a struct (namely, c)
	fmt.Println(d.z)    // d does not need to be dereferenced to access d.z
	fmt.Println((*d).z) // or we could go through d the hard way
}

// C is a simple struct holding a single value z.
type C struct {
	z int
}

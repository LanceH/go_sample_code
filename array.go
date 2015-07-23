package main

/*
From "Effective Go"

    - Arrays are values. Assigning one array to another copies all the elements.
    - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
    - The size of an array is part of its type. The types [10]int and [20]int are distinct.

Arrays are very strict and limited and are used internally to build slices.  See slices.go.

They are still useful for exact control of memory.
*/

import (
	"fmt"
)

func main() {
	// array initialized to values that the base type would initialized to
	var a [3]int64
	for i, j := range a {
		fmt.Printf("index: %d - value: %d\n\n", i, j)
	}

	// values are copied during assignment -- b doesn't point to the same memory as a
	b := a
	b[0] = 1
	b[1] = 2
	b[2] = 3
	fmt.Printf("a[2]: %d - b[2]: %d\n\n", a[2], b[2])

	// passing an array to a function copies it
	// changes in that function don't affect the original
	copyTest(a)
	fmt.Printf("a[0] is unchanged: %d\n\n", a[0])

	// [3]int64 is an array not a slice
	// append(a,4)
	// will not compile:
	// "first argument to append must be slice; have [3]int64"

	// This is a slice
	var s []int64
	fmt.Println(s)
	s = append(s,1)
	fmt.Println(s)
}

// the size of the array is part of the type and must be specified
func copyTest(c [3]int64) {
	c[0] = 186000
}

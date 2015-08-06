package main

import "fmt"

/*
From "Effective Go":

  Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.

  Slices are generally the idiomatic way of doing stuff normally done in arrays in other languages.
*/

var a []int64

func main() {
	// a is a slice initialized as nil: []
	// the capacity of the slice is doubled when it needs more
	// room in the underlying array
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n", a, len(a), cap(a))
	a = append(a, 55)
	fmt.Printf("%s has length: %d and capacity %d\n\n", a, len(a), cap(a))

	// size of a slice can be described upon initialization using make()
	b := make([]int64, 100, 1000)
	fmt.Printf("%d has length: %d and capacity %d\n\n", b, len(b), cap(b))

	// make an array
	// then make a slice reference the exact same array
	array := [3]int64{1, 2, 3}
	slice := array[:]
	fmt.Printf("%d has length: %d and capacity %d\n\n", slice, len(slice), cap(slice))
	// this changes the underlying array
	slice[2] = 12
	fmt.Printf("%d has length: %d and capacity %d\n\n", slice, len(slice), cap(slice))

	// once we append and increase the capacity, it now points to a different underlying array
	slice = append(slice, 4)
	fmt.Printf("%d has length: %d and capacity %d\n\n", slice, len(slice), cap(slice))
	fmt.Println(array)
}

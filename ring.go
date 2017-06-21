package main

/*
container/ring is kind of like a linked list, but it keeps going around.

Explores how rings are merged.
*/
import (
	"container/ring"
	"fmt"
)

// Pos describes coordinates on the xy plane.
type Pos struct {
	s string
}

func main() {
	p1 := Pos{s: "first"}
	p2 := Pos{s: "second"}
	p3 := Pos{s: "third"}
	p4 := Pos{s: "fourth"}

	r := ring.New(1)
	r.Value = p1 // At the first element
	fmt.Printf("%+v\n", r.Value)

	r2 := ring.New(1)
	r2.Value = p2
	r.Link(r2)   // still at first element
	r = r.Next() // at the second element
	fmt.Printf("%+v\n", r.Value)

	r3 := ring.New(1)
	r3.Value = p3
	r.Link(r3)   // still at second element
	r = r.Next() // at the third element

	r4 := ring.New(1)
	r4.Value = p4
	r.Link(r4)    // still at third element
	r = r.Move(2) // move back to the first

	// Iterate through ring and print its contents.
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value.(Pos).s)
		r = r.Next()
	}

	fmt.Println("Next element is: ", r.Next().Value.(Pos).s)
	fmt.Println("Removing next element...")
	r.Unlink(1)
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value.(Pos).s)
		r = r.Next()
	}
}

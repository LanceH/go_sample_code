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
	x float64
	y float64
}

func main() {
	p1 := Pos{x: 1.1, y: 1.1}
	p2 := Pos{x: 2.2, y: 2.2}
	p3 := Pos{x: 3.3, y: 3.3}
	p4 := Pos{x: 4.4, y: 4.4}

	r := ring.New(1)
	r.Value = p1
	r2 := ring.New(1)
	r2.Value = p2
	r.Link(r2)
	r3 := ring.New(2)
	r3.Value = p3
	r3.Next().Value = p4
	r3 = r3.Next()
	r3.Link(r)
	r4 := r3.Next()
	fmt.Println("r3:", r3.Value.(Pos).x, "r4", r4.Value.(Pos).x)

	// Iterate through ring and print its contents.
	for i := 0; i < r.Len()*2; i++ {
		fmt.Println(r.Value.(Pos).x)
		fmt.Println(r.Value.(Pos).y)
		r = r.Next()
	}

}

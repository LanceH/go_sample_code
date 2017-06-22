package main

import (
	"container/ring"
	"fmt"
)

type rg struct {
	*ring.Ring
}

var r rg

func main() {

	r = rg{ring.New(1)}
	r.Value = "blah"
	r.doit()
}

func (r rg) doit() {
	fmt.Println(r.Value)
}

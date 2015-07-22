package main

/*
import fmt directly to main (no 'fmt' prefix)
alias the math package to 'm'

This will result in a golint error for the . import.  It is not
recommended to do a dot import because it may cause conflicts with
future versions of go.
*/

import (
	. "fmt"
	m "math"
)

func main() {
	Println(m.Pi)
}

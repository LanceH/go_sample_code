package main

/*
Run from this directory:
go tool vet .

The fmt.Printf statement below is written to provide
a go vet warning due to a mismatch of %s asking for a
string and 33 being an int.
*/

import (
	"fmt"
)

func main() {
	fmt.Printf("%s", 33)
}

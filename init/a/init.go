package a

import (
	"../c"
)

func init() {
	c.X("init() in package a")
}

// X calls c.X
func X() {
	c.X("A")
}

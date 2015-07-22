package b

import (
	"../c"
)

func init() {
	c.X("init() in package b")
}

// X calls c.X
func X() {
	c.X("B")
}

package b

import (
	"../c"
)

func init() {
	c.X("init() in package b")
}

func X() {
	c.X("B")
}

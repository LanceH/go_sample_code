package a

import (
	"../c"
)

func init() {
	c.X("init() in package a")
}

func X() {
	c.X("A")
}

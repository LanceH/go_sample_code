package main

import ("fmt")

var m = make(map[string] string)

func main() {
	fmt.Println(m) // nil map
	m["a"] = "A"
	fmt.Println(m)
	fmt.Println(m["a"])
}
package main

/*
Create and use a Struct.
*/

import ("fmt")

type Car struct {
  Speed int
}

func main() {
  c1 := Car{44}
  fmt.Println(c1.Speed)
  c2 := Car{Speed: 55}
  fmt.Println(c2.Speed)
}

package main

import ("fmt")

func main() {
  a := 77
  b := &a
  fmt.Println(a)  // a is a normal int value
  fmt.Println(b)  // b is a pointer to a (0x20818a220 or something like that)
  c := C{c: 77}
  fmt.Println(c)   // c is a struct
  fmt.Println(c.c) // c.c is the value 77 again
  d := &c
  fmt.Println(d)   // d is a pointer to a struct
  fmt.Println(d.c) // d is a pointer but go figures it out and keeps going
  fmt.Println((*d).c) // or we could go through d the hard way
}

type C struct {
  c int
}
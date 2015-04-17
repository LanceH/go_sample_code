package main

/*
init() is a special function called before main.

I managed to read two books before learning this one.
*/

import(
  "fmt"
  )

func main() {
  fmt.Println("Main")
}

func init() {
  fmt.Println("Init")
}
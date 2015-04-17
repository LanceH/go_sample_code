package main

/* 
Pi  = 3.14159265358979323846264338327950288419716939937510582097494459
*/

import (
  "fmt"
  "math"
)

func main() {
  a := math.Pi
  fmt.Printf("%f\n", a)
  a = math.Pi - 3.141593
  fmt.Printf("%f\n", a)
  fmt.Println(a)
  fmt.Println(math.Pi)
}
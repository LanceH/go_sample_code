package main

/*
import fmt directly to main (no 'fmt' prefix)
alias the math package to 'm'
*/

import (
  . "fmt"
  m "math"
)

func main() {
  Println(m.Pi)
}

package main

import ("fmt")

func main() {
  a := []int{1,2,3,4}
  fmt.Println(Sum(a))
}

func Sum(a []int) (s int){
  for _, i := range a {
    s += i
  }
  return s
}
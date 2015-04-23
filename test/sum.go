package main

import ("fmt")

func main() {
  fmt.Println([4]int{1,2,3,4})
}

func Sum(a []int) (s int){
  for _, i := range a {
    s += i
  }
  return s
}

func Average(a []int) (s float64){
  for _, i := range a {
    s += float64(i)
  }
  return s / float64(len(a))
}
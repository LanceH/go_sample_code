package main

/*
go test -v
*/

import "testing"

func TestSum(t *testing.T) {
  v := Sum([]int{1,2,3,4})
  if v!= 10 {
    t.Error("Expected 10, got ", v)
  }
}

func TestAverage(t *testing.T) {
  v := Average([]int{1,2,3,4})
  if v!= 2.5 {
    t.Error("Expected 2.5, got ", v)
  }
}

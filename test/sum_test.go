package main

import "testing"

func TestSum(t *testing.T) {
  v := Sum([]int{1,2,3,4})
  if v!= 10 {
    t.Error("Expected 10, got ", v)
  }
}
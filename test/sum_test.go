package main

/*
go test -v
go test -bench .
*/

import "testing"

var list []int = []int{1,2,3,4}


func TestSum(t *testing.T) {
  v := Sum(list)
  if v!= 10 {
    t.Error("Expected 10, got ", v)
  }
}

func TestAverage(t *testing.T) {
  v := Average(list)
  if v!= 2.5 {
    t.Error("Expected 2.5, got ", v)
  }
}

func BenchmarkSum(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Sum(list) //list come from sum_test.go
  }
}

func BenchmarkAverage(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Average(list) //list come from sum_test.go
  }
}

func BenchmarkNothing(b *testing.B) {
  for i := 0; i < b.N; i++ {
  }
}
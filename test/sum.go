package main

import (
	"fmt"
)

func main() {
	fmt.Println([4]int{1, 2, 3, 4})
}

/*
Sum returns the sum of all the elements of an array.
*/
func Sum(a []int) (s int) {
	for _, i := range a {
		s += i
	}
	return s
}

/*
Average returns the average of all the elements of an array.
*/
func Average(a []int) (s float64) {
	for _, i := range a {
		s += float64(i)
	}
	return s / float64(len(a))
}

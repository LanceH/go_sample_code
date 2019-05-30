package main

import "fmt"

var fibs = make(map[int]int)

func main() {
	var i int

	fmt.Print("Integer > 0: ")
	for _, err := fmt.Scanf("%d", &i); err != nil; {
	}
	fmt.Println(fib(i))
}

func fib(i int) (n int) {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 1
	} else if fibs[i] != 0 {
		return fibs[i]
	}
	n = fib(i-1) + fib(i-2)
	fibs[i] = n
	return n
}

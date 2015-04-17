package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	for i := 1; i <= 10; i++ {
		<-ticker
		fmt.Printf("\rOn %d/10", i)
	}
	fmt.Println("\nAll is said and done.")
}

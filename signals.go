package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Starting in Main()")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	s := <-c
	fmt.Println("Received signal: ", s)
}

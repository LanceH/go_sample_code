package main

import (
	"log"
	"os"
)

func main() {
	// Eat the error with _
	file, _ := os.Open("errors.go")
	defer file.Close()

	// The normal way of handling errors
	file, err := os.Open("./tmp/new.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

}

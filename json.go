package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

var a = `{"one": "1"}`

func main() {

	decoder := json.NewDecoder(strings.NewReader(a))
	for {
		var n Number
		if err := decoder.Decode(&n); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("one: %s\n", n.One)
	}
}

type Number struct {
	One string
}

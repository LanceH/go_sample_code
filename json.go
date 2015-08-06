package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

var a = `{"one": "1"}`
var b = `{"C": {"x": "3"}, "one" : "1"}`

func main() {

	decoder := json.NewDecoder(strings.NewReader(a))
	for {
		var n number
		if err := decoder.Decode(&n); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("one: %s\n", n.One)
	}

	decoder = json.NewDecoder(strings.NewReader(b))
	for {
		var b bee
		if err := decoder.Decode(&b); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	}
}

type number struct {
	One string
}

type bee struct {
	C   cee
	One string
}

type cee struct {
	X string
}

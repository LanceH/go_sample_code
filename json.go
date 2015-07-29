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
		var n Number
		if err := decoder.Decode(&n); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("one: %s\n", n.One)
	}

	decoder = json.NewDecoder(strings.NewReader(b))
	for {
		var b Bee
		if err := decoder.Decode(&b); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	}
}

type Number struct {
	One string
}

type Bee struct {
	C   *Cee
	One string
}

type Cee struct {
	X string
}

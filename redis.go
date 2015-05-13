package main

import (
	"fmt"
	"redis"
)

func main() {
	spec := redis.DefaultSpec().Password("foobared")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		fmt.Println("failed to connect ", e)
	}
	for i := 0; i < 73; i++ {
		a, _ := client.Incr("a")

		fmt.Println(a)
	}
}

package main

/* this program has some weirdness
because count is both read and written to
by different threads, potentially as the same time
causing a data race -- bad

If a data race doesn't happen, then it should print out
count every second
*/


import (
  "fmt"
  "time"
  "runtime"
)

var count int64 = 0

func main() {
  runtime.GOMAXPROCS(2)

  fmt.Println("main")
  go Counter()
  fmt.Println("after Counter()")
  for {
    count++
  }
}

func Counter() {
  fmt.Println("In Counter()")
  for {
    fmt.Println(count)
    time.Sleep(time.Second)
  }
}

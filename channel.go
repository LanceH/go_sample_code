package main

/*
Channels are built in -- no package required

They block until they match up.
*/


import(
  "fmt"
  "time"
)


func main() {
  a := make(chan int)
  go handle(a) // make a go routine to wait around to do something with stuff put into a
  a <- 7 // put something into the channel
  time.Sleep(time.Second * 1) // this is here to make sure main() doesn't finish before handle()
}

func handle(a chan int) {
  fmt.Println(<-a)  // this go routine blocks until something is put into the channel
}

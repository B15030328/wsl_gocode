package main

import "fmt"

func main() {
	var chan1 = make(chan int, 2)

	go func() {
		fmt.Println("hello world")
		chan1 <- 1
	}()

	<-chan1
}

package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int)

	go func() {
		fmt.Println("channel test")
		chan1 <- 12
	}()
	fmt.Println(<-chan1)
}

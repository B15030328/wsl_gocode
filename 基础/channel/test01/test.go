package main

import (
	"fmt"
	"time"
)

func main() {
	stopc := make(chan struct{})
	select {
	case <-time.After(100 * time.Millisecond):
		close(stopc)
	case <-stopc:
		fmt.Println("11")
	}
}

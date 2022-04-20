package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomic_add(wg *sync.WaitGroup, total *uint64) {
	defer wg.Done()
	var i uint64
	for i = 1; i < 10; i++ {
		atomic.AddUint64(total, i)
	}

}

func main() {
	var wg sync.WaitGroup
	var total uint64
	wg.Add(1)
	go atomic_add(&wg, &total)
	wg.Wait()
	fmt.Println(total)

}

package main

import (
	"sync"
	"testing"
)

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B)    { benchmark(b, &mutex{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmark(b, &RWmutex{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &mutex{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWmutex{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmark(b, &mutex{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmark(b, &RWmutex{}, 5, 5) }

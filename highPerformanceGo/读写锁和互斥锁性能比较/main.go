package main

import (
	"sync"
	"time"
)

type RW interface {
	Read()
	Write()
}

const cost = time.Millisecond

type mutex struct {
	count int
	mu    sync.Mutex
}

func (m *mutex) Read() {
	m.mu.Lock()
	defer m.mu.Unlock()
	_ = m.count
	time.Sleep(cost)
}

func (m *mutex) Write() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count++
	time.Sleep(cost)
}

type RWmutex struct {
	count int
	mu    sync.RWMutex
}

func (rw *RWmutex) Read() {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	_ = rw.count
	time.Sleep(cost)
}

func (rw *RWmutex) Write() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.count++
	time.Sleep(cost)
}

func main() {

}

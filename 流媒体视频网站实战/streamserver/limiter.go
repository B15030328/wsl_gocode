package main

import "log"

/*
利用channel实现连接流量控制
share channel instead of share memory
*/

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) getConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("connection reach the limit")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) releaseConn() {
	c := <-cl.bucket
	log.Println("new connection comes, ", c)
}

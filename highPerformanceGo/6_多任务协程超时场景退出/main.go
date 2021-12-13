package main

import (
	"fmt"
	"time"
)

/*
场景：在协程需要执行两个任务时，判断第一个任务超时就退出
*/
func multiTask(phrase, done chan bool) {
	time.Sleep(time.Second) //模拟允许的第一个任务执行时间
	select {
	case phrase <- true:
	default:
		return
	}
	time.Sleep(time.Second) //模拟允许的第二个任务执行时间
	done <- true
}

func main() {
	phrase := make(chan bool)
	done := make(chan bool)
	go multiTask(phrase, done)
	select {
	case <-phrase:
		fmt.Println("第一个任务完成")
		<-done
		fmt.Println("第二个任务完成")
	case <-time.After(time.Second * 2):
		fmt.Println("任务超时")
	}
}

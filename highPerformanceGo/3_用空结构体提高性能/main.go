package main

/*
空结构体不占用字节
妙用：
1、实现set
2、当channel不传输数据，只用来传递执行信号时用struct{}{}作为占位符
*/

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("空结构体占用字节数:", unsafe.Sizeof(struct{}{}))

	var i int64 = 1

	fmt.Println("int64占用字节数:", unsafe.Sizeof(i))

	channel_use()

}

// 2、channel妙用
func channel_use() {
	fmt.Println("测试空结构体在channel中的妙用")
	chan1 := make(chan struct{})
	go work(chan1)
	time.Sleep(time.Second)
	chan1 <- struct{}{}
	time.Sleep(time.Second)
}

func work(chan1 chan struct{}) {
	<-chan1
	fmt.Println("开始执行任务")
	close(chan1)
}

package main

//https://www.qcrao.com/2019/06/12/dive-into-go-context/

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute/2)

	//取消goroutine示例
	go print(ctx)
	time.Sleep(time.Minute)
	cancle()

	//取消goroutine示例(将上面注释)
	// defer cancle()
	// for num := range gen(ctx) {
	// 	fmt.Println(num)
	// 	if num == 5 {
	// 		cancle()
	// 		break
	// 	}
	// }

}

//取消goroutine示例
func print(ctx context.Context) {
	i := 0
	for {
		fmt.Printf("第%d秒啦\n", i)
		i++

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}

//防止goroutine泄露
func gen(ctx context.Context) chan int {
	chan1 := make(chan int)
	go func() {
		var n = 0
		for {
			select {
			case <-ctx.Done():
				return
			case chan1 <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return chan1
}

package main

import (
	"fmt"
	"strings"
)

// 一般推荐使用 strings.Builder 来拼接字符串(倍数申请新内存)

func main() {
	var str1 = "hello"
	var builder strings.Builder
	cap := 0
	for i := 0; i < 1000; i++ {
		if builder.Cap() != cap {
			fmt.Println(builder.Cap())
			cap = builder.Cap()
		}
		builder.WriteString(str1)
	}
	fmt.Println(builder.String())
}

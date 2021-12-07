package main

// https://qcrao.com/2019/06/03/dive-into-go-unsafe/

import (
	"fmt"
	"unsafe"
)

type user struct {
	name  string
	age   int
	grade string
}

func main() {
	var a int32 = 2
	fmt.Println(unsafe.Pointer(&a), unsafe.Sizeof(a))

	user1 := user{name: "chuyu", age: 18, grade: "研三"}
	fmt.Println(user1)
	// 使用场景
	// 获取指针位置来修改值，获取偏移量来修改值
	lo1 := (*string)(unsafe.Pointer(&user1))
	*lo1 = "chory"
	fmt.Println(user1)

	lo2 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&user1)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lo2 = "24"
	fmt.Println(user1)
	//slice和map的相互转话
}

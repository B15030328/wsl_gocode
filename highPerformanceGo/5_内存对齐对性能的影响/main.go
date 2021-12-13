package main

import (
	"fmt"
	"unsafe"
)

type NoAlign struct {
	age    int64
	gender int32
}

type Align struct {
	age    int64
	gender int64
}

func main() {
	noalign := NoAlign{}
	align := Align{}
	fmt.Println("noalign", unsafe.Sizeof(noalign))
	fmt.Println("align", unsafe.Sizeof(align))
}

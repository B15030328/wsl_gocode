package main

import "fmt"

func foo() *int {
	t := 10
	return &t
}

func main() {
	x := foo()
	fmt.Println(*x)
}

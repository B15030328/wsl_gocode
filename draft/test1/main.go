package main

import "fmt"

func init() {
	fmt.Println("测试 init")
}
func main() {
	a := []int{1, 2, 3}
	// var a = make([]int, 1, 10)
	a = append(a, 1)
	fmt.Println(a)
	a = app(a)
	fmt.Println(a)

}

func app(a []int) []int {
	for i := 1; i < 4; i++ {
		a = append(a, i)
	}
	return a
}

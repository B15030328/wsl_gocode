package main

/*slice由于不会指向新的内存地址，所以原数组内存会一直被占用，当原数组占用内存较大时非常影响性能
这种情况可以用copy取切片而不是直接取
*/

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

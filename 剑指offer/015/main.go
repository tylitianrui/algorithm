package main

import "fmt"

func Nums(i int) int {
	var n int
	for i != 0 {
		// 取消二进制中最右侧的1
		i = i & (i - 1)
		n++
	}
	return n

}
func main() {
	fmt.Println(Nums(7))

}

package main

func hammingWeight(num uint32) int {
	var n int
	for num != 0 {
		// 取消二进制中最右侧的1
		num = num & (num - 1)
		n++
	}
	return n
}

package main

import "fmt"

/*
在一个数组中，只有一个数字出现一次（或奇数次），其他数字均出现两次（或偶数次），找出那个只出现一次（或奇数次）的数字

思路：采用异或
说明：异或。两个相同数字进行异或=0，0和A进行异或=A
将所有数字进行异或，即可获取这个数
*/
func FindSingle(l []int) int {
	s := 0
	for _, value := range l {
		s = s ^ value
	}
	return s
}
func main() {

	l := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	single := FindSingle(l)
	fmt.Println(single)
}

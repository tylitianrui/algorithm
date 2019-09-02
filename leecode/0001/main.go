package main

import "fmt"

// 代码
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	l := make([]int, 2)
	for idx, value := range nums {
		if v, ok := numMap[value]; ok {
			l[0] = v
			l[1] = idx
			break
		} else {
			numMap[target-value] = idx
		}

	}
	return l

}

func main() {
	l := []int{1, 6, 2, 4, 5, 0, 23, 12}
	fmt.Println(twoSum(l, 3))

}

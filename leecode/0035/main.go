package main

import "fmt"

func searchInsert(nums []int, target int) int {

	var idx int
	for i := 0; i < len(nums); i++ {
		if target > nums[i] {
			idx++
		} else {
			break
		}
	}
	return idx

}

func main() {
	l := []int{1, 2, 6, 9}
	fmt.Println(searchInsert(l, 0))
}

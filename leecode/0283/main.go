package main

import "fmt"

func main() {

	n := []int{0, 1, 0, 3, 12}
	moveZeroes(n)
	fmt.Println(n)
}

// answer

func moveZeroes(nums []int) {
	var s int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[i], nums[s] = nums[s], nums[i]
		s++

	}

}

package main

func runningSum1(nums []int) []int {
	s := 0
	sum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		sum[i] = s
	}
	return sum

}

//  更优的解答 开辟新空间  空间复杂度O(1)
func runningSum(nums []int) []int {
	s := 0
	// sum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		nums[i] = s
	}
	return nums

}

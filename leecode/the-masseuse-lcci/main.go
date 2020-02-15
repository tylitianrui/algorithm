package main

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(massage(nums[:len(nums)-2])+nums[len(nums)-1], massage(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return max(rob(nums[:len(nums)-2])+nums[len(nums)-1], rob(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}

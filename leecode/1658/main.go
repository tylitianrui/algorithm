package main

/*
此题精妙之处在于如何把未知未见过的问题转换为已知见过的问题。
其实如果求是否存在两边的数和为x，等价于求是否有连续区间的总和等于sum - x（假设sum为所有元素的总和）。
如果存在这样的区间，那么其两边的元素和必然为x。如果存在多个这样的区间，取最长的那个区间len，则N-len就是我们要求的最少删除元素个数令x减少为0.
*/
func minOperations(nums []int, x int) int {
	var (
		sum, l int
		winSum int
		t      = -1
	)

	// 计算总和
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 反向计算
	target := sum - x
	if target == 0 {
		return len(nums)
	}

	// 滑动窗口 l  r  t记录最大连续的的长度
	// 每次向右滑动
	for r := 0; r < len(nums); r++ {
		winSum += nums[r]

		// 如果此时大于了target,那么l也向右滑动
		for l < r && winSum > target {
			winSum -= nums[l]
			l++
		}
		if winSum == target {
			// 判断是否大于之前的记录,进行更新
			if t < r-l+1 {
				t = r - l + 1
			}

		}

	}
	// 此时为找到
	if t == -1 {
		return -1
	}
	return len(nums) - t

}

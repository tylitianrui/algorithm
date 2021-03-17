package main

/*
 滑动窗口

*/
func lengthOfLongestSubstring(s string) int {
	var (
		sum  int
		left int
	)
	//  如果小于1 则为长度
	if len(s) <= 1 {
		return len(s)
	}
	// right边界 每次向右移动
	// 如果遇到相同的元素，left边界向右
	for right := 1; right < len(s); right++ {
		//  校验是否相等
		for i := left; i < right; i++ {
			if s[right] == s[i] {
				left = i + 1
				break
			}
		}
		// 计算长度
		if right-left+1 > sum {
			sum = right - left + 1
		}
	}
	return sum

}

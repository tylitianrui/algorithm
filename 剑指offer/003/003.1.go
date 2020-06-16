package main

// 题目一  长度为n的数组，元素值范围0--n-1，找出其中任意一个重复的数字
func FindOneRepetition(l []int) (int, bool) {

	for i := 0; i < len(l); i++ {
		for {
			if i == l[i] {
				break
			}
			if l[l[i]] == l[i] {
				return l[i], true
			}
			l[l[i]], l[i] = l[i], l[l[i]]
		}
	}
	return 0, false

}

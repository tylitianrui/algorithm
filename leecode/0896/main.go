package main

func isMonotonic(A []int) bool {
	var desc, incr int
	if len(A) <= 2 {
		return true
	}
	for i := 1; i < len(A); i++ {

		if A[i-1] < A[i] { // 前面比后面小，incr=1
			incr = 1
		} else if A[i-1] > A[i] { // 前面比后面大，desc=1
			desc = 1
		}
	}
	if (desc + incr) == 2 { //  desc, incr 都为1 则不是
		return false
	}

	return true
}

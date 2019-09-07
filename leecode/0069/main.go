package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	return BSearch(0, x, x)
}

func BSearch(l, r, n int) int {

	mid := (l + r) >> 1
	if mid*mid > n {
		return BSearch(l, mid-1, n)
	} else if (mid+1)*(mid+1) <= n {
		return BSearch(mid+1, r, n)
	} else {
		return mid
	}

}
func main() {

	mySqrt(8)
}

package main

func isPowerOfFour(num int) bool {
	return (num-1)&num == 0 && num%3 == 1
}

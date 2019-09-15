package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	p := x
	s := 0
	for x != 0 {
		s = x%10 + s*10
		x /= 10

	}
	if s == p {
		return true
	}
	return false

}

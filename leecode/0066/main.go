package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 0
	num := 1
	for i := len(digits) - 1; i >= 0; i-- {
		flag := (digits[i] + num + carry) / 10

		if flag == 0 {
			digits[i] = digits[i] + num + carry
			carry = 0
			break
		} else {

			digits[i] = (digits[i] + num + carry) % 10
			num = 0
			carry = flag

		}

	}
	fmt.Println(carry)
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{8, 9, 9, 9, 9}))
}

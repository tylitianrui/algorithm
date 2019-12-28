package main

func singleNumber(nums []int) int {
	res := 0
	for _, value := range nums {
		res ^= value
	}
	return res
}

func main() {
	l := []int{1, 2, 3, 4, 3, 2, 1}
	singleNumber(l)
}

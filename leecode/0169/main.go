package main

func main() {

}

func majorityElement(nums []int) int {
	var cnt, res int
	for _, value := range nums {
		if cnt == 0 {
			res = value
			cnt++
		} else if res == value {
			cnt++
		} else {
			cnt--
		}
	}
	return res

}

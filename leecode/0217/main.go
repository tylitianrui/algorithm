package main

func containsDuplicate(nums []int) bool {
	item := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := item[nums[i]]; ok {
			return true
		} else {
			item[nums[i]] = 0
		}

	}
	return false
}

func main() {

}

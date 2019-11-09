package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for {
		if l > r {
			break
		}

		mid := (l + r) / 2
		if nums[l] <= nums[mid] {
			if nums[mid] <= nums[r] {
				return nums[l]
			} else {
				l = mid + 1
			}

		} else {
			if nums[mid] <= nums[r] {
				r = mid
			}
		}

	}
	return -1

}
func main() {

}

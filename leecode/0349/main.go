package main

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	tmp := make(map[int]int)
	for _, value := range nums1 {
		tmp[value] = 0
	}

	for _, value := range nums2 {
		if k, ok := tmp[value]; ok && k == 0 {
			res = append(res, value)
			tmp[value] = 1
		}

	}
	return res
}

func main() {

}

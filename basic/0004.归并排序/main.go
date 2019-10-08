package main

import "fmt"

func merge(left, rigth []int) []int {
	i, j := 0, 0
	var res = make([]int, 0)
	for {

		if i == len(left) {
			res = append(res, rigth[j:]...)
			return res
		}
		if j == len(rigth) {
			res = append(res, left[i:]...)
			return res
		}

		if left[i] <= rigth[j] {
			res = append(res, left[i])
			i++
			continue
		}
		if left[i] > rigth[j] {
			res = append(res, rigth[j])
			j++
			continue
		}

	}

}

func mersort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	l := mersort(arr[:len(arr)/2])
	r := mersort(arr[len(arr)/2:])
	return merge(l, r)

}

func main() {
	l := []int{3, 1, 3, 54, 6, 2, 9, 45, 4, 56}
	fmt.Println(mersort(l))

}

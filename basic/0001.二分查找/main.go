package main

import "fmt"

func BS(arr []int, l, r, n int) int {
	if l > r {
		return -1
	}
	mid := (r + l) >> 1
	if arr[mid] > n {
		return BS(arr, l, mid-1, n)
	} else if arr[mid] == n {
		return mid

	} else {
		return BS(arr, mid+1, r, n)
	}

}

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(BS(l, 0, len(l)-1, 19))

}

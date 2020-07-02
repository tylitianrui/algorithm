package main

import "fmt"

func main() {
	l := []int{1, 6, 3, 9, 1, 0, 12, 9}
	InsertSort(l)
	fmt.Println(l)

}

func InsertSort(l []int) {
	if len(l) <= 0 {
		return
	}
	for i := 1; i < len(l); i++ {
		for j := i; j > 0; j-- {
			if l[j] < l[j-1] {
				l[j], l[j-1] = l[j-1], l[j]

			}

		}

	}

}

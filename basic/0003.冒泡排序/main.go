package main

import "fmt"

func BS(arr []int) {
	for i := 0; i < len(arr); i++ {
		order := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				order = false
			}

		}
		if order {
			return

		}
	}
}
func main() {
	l := []int{1, 4, 28, 32, 12, 55, 10}
	BS(l)
	fmt.Println(l)

}

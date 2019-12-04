package main

import "fmt"

func trap(height []int) int {
	var res int
	if len(height) < 3 {
		return res
	}

	stack := make([]int, 0)

	for i := 0; i < len(height); i++ {

		for {

			if len(stack) == 0 || height[stack[len(stack)-1]] >= height[i] {
				break

			}

			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) != 0 {

				res += (min(height[stack[len(stack)-1]], height[i]) - height[tmp]) * (i - stack[len(stack)-1] - 1)

			}

		}

		stack = append(stack, i)

	}
	return res

}

func min(a, b int) int {
	if a > b {
		return b

	}
	return a
}
func main() {
	l := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	r := trap(l)
	fmt.Println(r)

}

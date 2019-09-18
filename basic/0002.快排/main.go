package main

import "fmt"

func Partion(l []int, left, right int) int {
	x := l[right]
	i := left - 1
	for j := left; j < right; j++ {
		if l[j] <= x {
			i++
			l[i], l[j] = l[j], l[i]
		}
	}
	l[i+1], l[right] = l[right], l[i+1]
	return i + 1
}

func qs(l []int, left, right int) {
	if left < right {
		p := Partion(l, left, right)
		qs(l, left, p-1)
		qs(l, p+1, right)
	}

}
func main() {
	l := []int{2, 7, 12, 6, 19, 0, -1, 4}
	qs(l, 0, len(l)-1)
	//Partion(l,0,len(l)-1)
	fmt.Println(l)

}

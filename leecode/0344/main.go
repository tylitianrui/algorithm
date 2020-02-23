package main

import "fmt"

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	l := 0
	r := len(s) - 1
	for {
		if l >= r {
			return
		}
		s[l], s[r] = s[r], s[l]
		l++
		r--

	}

}
func main() {
	l := []byte("")
	reverseString(l)
	fmt.Println(string(l))

}

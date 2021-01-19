package main

import (
	"errors"
	"fmt"
)

var target = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func main() {
	i := F16to10("a3456")
	fmt.Println(i)

}

func F16to10(n string) int {

	length := len(n)
	if length == 1 {
		if idx, err := index(n, target); err != nil {
			panic(err.Error())
		} else {
			return idx
		}

	}
	return F16to10(n[:length-1])*16 + F16to10(n[length-1:])
}

func index(o string, li []string) (int, error) {
	for i := 0; i < len(li); i++ {
		if o == li[i] {
			return i, nil
		}
	}
	return -1, errors.New("")

}

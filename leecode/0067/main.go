package main

import "fmt"

func addBinary(a string, b string) string {
	var (
		// 0 - 48  ; 1 - 49
		carry  int = 48
		length int
		res    string = ""
		l      string = ""
	)

	if len(a) >= len(b) {
		length = len(b)
		l = a[:len(a)-length]

	} else {
		length = len(a)
		l = b[:len(b)-length]
	}
	//fmt.Println(l)

	for i := 0; i < length; i++ {
		ca := a[len(a)-i-1]
		cb := b[len(b)-i-1]
		//fmt.Println(ca)
		//fmt.Println(cb)
		if ca == 49 && cb == 49 {
			if carry == 48 {
				res = "0" + res
			} else {
				res = "1" + res
			}
			carry = 49

		} else if ca == 48 && cb == 48 {
			if carry == 49 {
				res = "1" + res
			} else {
				res = "0" + res
			}
			carry = 48
		} else {
			if carry == 49 {
				res = "0" + res
				carry = 49
			} else {
				res = "1" + res
				carry = 48
			}

		}

	}
	fmt.Println(carry)

	for i := 0; i < len(l); i++ {
		if carry == 48 {
			fmt.Println(l[:len(l)-1-i])
			res = l[:len(l)-i] + res
			break
		} else {
			c := l[len(l)-1-i]
			if c == 48 {
				res = "1" + res
				carry = 48
			} else {
				res = "0" + res
				carry = 49
			}
		}

	}
	if carry == 49 {
		res = "1" + res
	}

	return res

}

func main() {
	t := addBinary("100", "110010")
	fmt.Println(t)

}

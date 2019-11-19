package main

import "fmt"

func convertToBase7(num int) string {
	if num < 0 {
		num = 0 - num
		return fmt.Sprintf("-%s", a(num))
	}
	return a(num)

}

func a(num int) string {
	if num < 7 {
		return fmt.Sprintf("%d", num)
	}
	return fmt.Sprintf("%s%s", convertToBase7(num/7), convertToBase7(num%7))
}

func main() {

}

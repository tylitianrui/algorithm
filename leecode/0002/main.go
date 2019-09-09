package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	n := p
	carry, sum, a, b := 0, 0, 0, 0
	for {
		if l2 == nil && l1 == nil {
			break
		}
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val

		}

		if l2 == nil {
			b = 0
		} else {
			b = l2.Val

		}
		sum, carry = add(a, b, carry)
		n.Next = &ListNode{Val: sum, Next: nil}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		n = n.Next

	}
	if carry == 1 {
		n.Next = &ListNode{Val: 1, Next: nil}

	}
	return p.Next
}

func add(a, b, c int) (int, int) {
	m := a + b + c
	return m % 10, m / 10
}

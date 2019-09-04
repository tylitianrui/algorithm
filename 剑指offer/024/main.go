package main

type Node struct {
	Val  int
	Next *Node
}

// 链表翻转
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	node := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node

}

func main() {

}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	l := Mid(root)
	min := l[1] - l[0]
	for i := 0; i < len(l)-1; i++ {
		if l[i+1]-l[i] < min {
			min = l[i+1] - l[i]
		}
	}
	return min
}

func Mid(root *TreeNode) []int {
	mid := []int{}
	if root == nil {
		return mid
	}
	helper(root, &mid)
	return mid

}

func helper(root *TreeNode, mid *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, mid)
	*mid = append(*mid, root.Val)
	helper(root.Right, mid)

}

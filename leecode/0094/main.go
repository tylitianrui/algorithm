package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper(root, &res)
	return res
}

// 中序遍历 存入slice之中
func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, res)
	*res = append(*res, root.Val)
	helper(root.Right, res)
}

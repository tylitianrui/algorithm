package main

// 二叉查找树，有序！
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}

}

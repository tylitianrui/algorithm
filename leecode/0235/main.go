package main

// 二叉查找树，有序！
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 答案
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	} else if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}

}

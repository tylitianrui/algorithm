package _1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return j(root.Left, root.Right)

}

func j(left, rigth *TreeNode) bool {
	if left == nil && rigth == nil {
		return true
	}
	if left != nil && rigth != nil && left.Val == rigth.Val && j(left.Left, rigth.Right) && j(left.Right, rigth.Left) {
		return true
	}
	return false

}

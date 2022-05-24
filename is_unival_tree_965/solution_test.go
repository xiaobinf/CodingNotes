package is_unival_tree_965

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

func isUnivalTree(root *TreeNode) bool {
	num := root.Val
	return isUnival(root.Left, num) && isUnival(root.Right, num)
}

func isUnival(root *TreeNode, num int) bool {
	if root == nil {
		return true
	}
	if num != root.Val {
		return false
	}
	return isUnival(root.Left, num) && isUnival(root.Right, num)
}

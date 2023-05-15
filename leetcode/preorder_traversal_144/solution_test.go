package preorder_traversal_144

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

func preOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	preOrder(root.Left, arr)
	preOrder(root.Right, arr)
}

func preorderTraversal(root *TreeNode) []int {
	var arr []int
	preOrder(root, &arr)
	return arr
}

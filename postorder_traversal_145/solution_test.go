package postorder_traversal_145

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

func postOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, arr)
	postOrder(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	var arr []int
	postOrder(root, &arr)
	return arr
}

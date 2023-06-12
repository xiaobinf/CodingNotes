package inorder_traversal_94

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归左中右
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorder(root *TreeNode, res *[]int) {
//	if root == nil {
//		return
//	}
//	inorder(root.Left, res)
//	*res = append(*res, root.Val)
//	inorder(root.Right, res)
//}
//
//func inorderTraversal(root *TreeNode) []int {
//	var res []int
//	inorder(root, &res)
//	return res
//}

// 复习
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

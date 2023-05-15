package sum_of_left_leaves_404

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

// 错误做法 二叉树递归不应该 往下多层思考
//func sumOfLeftLeaves(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 0
//	}
//	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil && root.Right == nil {
//		return root.Left.Val
//	}
//	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil && root.Left == nil {
//		return 0
//	}
//	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil && root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
//		return root.Left.Val
//	}
//	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
//}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeft(root.Left, true) + sumOfLeft(root.Right, false)
}

func sumOfLeft(node *TreeNode, flag bool) int {
	if node == nil {
		return 0
	}
	// 判断如果是左叶子节点
	if node.Left == nil && node.Right == nil {
		if flag {
			return node.Val
		}
		return 0
	}
	return sumOfLeft(node.Left, true) + sumOfLeft(node.Right, false)
}

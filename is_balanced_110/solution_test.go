package is_balanced_110

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// todo 优化成自底向上的方法
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int16) int16 {
	if a < 0 {
		return -a
	}
	return a
}

// 比较数的高度
func max(a, b int16) int16 {
	if a < b {
		return b
	}
	return a
}

func getDepth(root *TreeNode) int16 {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(getDepth(root.Left)-getDepth(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

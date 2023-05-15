package is_symmetric_101

// 判断二叉树是否对称
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

func symmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else {
		return p.Val == q.Val && symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
	}
}

func isSymmetric(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}

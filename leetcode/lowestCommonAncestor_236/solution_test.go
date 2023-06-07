package lowestCommonAncestor_236

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor
// root为空，p，q为root时 直接返回root
//
// 递归左右，都不为空说明一个在左，一个在右 此时返回root
//
// 此时谁不为空返回谁 简洁写法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果root等于所求节点，那么root就是公共祖先
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 判断左右子树是否存在p，q的公共祖先。 会一直往下定位到p和q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右子树都存在
	if left != nil && right != nil {
		return root
	}

	// 如果左子树不存在，那就是右子树
	if left == nil {
		return right
	}

	return left
}

func TestName(t *testing.T) {

}

package diameterOfBinaryTree_543

import (
	"fmt"
	"testing"
)

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

// ans会更新每个节点的最大直径值
var ans = 0

func diameterOfBinaryTree(root *TreeNode) int {
	depth(root)
	return ans - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left)
	r := depth(node.Right)
	ans = max(ans, l+r-2)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestName(t *testing.T) {
	fmt.Println()
}

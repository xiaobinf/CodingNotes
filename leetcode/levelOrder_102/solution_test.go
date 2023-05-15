package levelOrder_102

import "testing"

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

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode //该切片用于存储二叉树每一层的节点
		for j := 0; j < len(q); j++ {
			// 遍历第i层的节点 并处理每个节点的左右节点
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 更新每一层需要遍历的队列
		q = p
	}
	return res
}

func TestName(t *testing.T) {

}

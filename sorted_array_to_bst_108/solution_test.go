package sorted_array_to_bst_108

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

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayConvertToBST(nums, 0, len(nums)-1)
}

// 本质上是找中间数作为中间节点，左右递归同理
// 确定结束条件
func sortedArrayConvertToBST(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	node.Left = sortedArrayConvertToBST(nums, start, mid-1)
	node.Right = sortedArrayConvertToBST(nums, mid+1, end)
	return node
}

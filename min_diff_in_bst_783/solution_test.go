package min_diff_in_bst_783

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

// 二叉搜索树节点之间的最小差值 其实就是排序数组元素间的最小差值

//func minDiffInBST(root *TreeNode) int {
//	// 中序遍历
//	var arr []int
//	inOrder(root, &arr)
//	var minNum = arr[1] - arr[0]
//	// 遍历切片
//	for i := 1; i < len(arr); i++ {
//		temp := arr[i] - arr[i-1]
//		if temp < minNum {
//			minNum = temp
//		}
//	}
//	return minNum
//}
//
//func inOrder(node *TreeNode, arr *[]int) {
//	if node == nil {
//		return
//	}
//	inOrder(node.Left, arr)
//	*arr = append(*arr, node.Val)
//	inOrder(node.Right, arr)
//}

// 解法二 可以不用存在数组里面 记录一个前置节点
var pre *TreeNode
var minN = int(^uint32(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDiffInBST(root *TreeNode) int {
	// 中序遍历
	inOrder(root)
	return minN
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	if pre != nil {
		minN = min(node.Val-pre.Val, minN)
	}
	pre = node
	inOrder(node.Right)
}

func TestName(t *testing.T) {
	fmt.Println(int(^uint32(0) >> 1))
}

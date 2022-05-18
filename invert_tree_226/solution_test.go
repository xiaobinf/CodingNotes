package invert_tree_226

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

// 自顶向下 递归 左右子树交换
func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return root
  }
  // 交换左右叶子节点
  var node *TreeNode
  node = root.Left
  root.Left = root.Right
  root.Right = node
  invertTree(root.Left)
  invertTree(root.Right)
  return root
}

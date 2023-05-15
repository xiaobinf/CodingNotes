package has_path_sum_112

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

// 只考虑正数场景
func hasPathSum1(root *TreeNode, targetSum int) bool {
  if root == nil {
    return false
  }
  if root.Val > targetSum {
    return false
  }
  if root.Val == targetSum {
    if root.Left == nil && root.Right == nil {
      return true
    }
  }
  if root.Left == nil && root.Right == nil {
    return false
  } else if root.Left == nil && root.Right.Val > targetSum-root.Val {
    return false
  } else if root.Right == nil && root.Left.Val > targetSum-root.Val {
    return false
  } else {
    return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
  }
}

// 还需要考虑负数场景
func hasPathSum(root *TreeNode, targetSum int) bool {
  if root == nil {
    return false
  }
  // 由于存在负数，0等等场景 不能直接判断大于 小于。
  if root.Val == targetSum {
    if root.Left == nil && root.Right == nil {
      return true
    }
    // 此处不能返回 false 有可能后面的节点和为0
  }
  if root.Left == nil && root.Right == nil {
    return false
  } else if root.Left == nil {
    return hasPathSum(root.Right, targetSum-root.Val)
  } else if root.Right == nil {
    return hasPathSum(root.Left, targetSum-root.Val)
  } else {
    return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
  }
}

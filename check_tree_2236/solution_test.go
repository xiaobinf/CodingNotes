package check_tree_2236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return (root.Left.Val + root.Right.Val) == root.Val
}

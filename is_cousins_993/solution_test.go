package is_cousins_993

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

// 思路整理
// 本质上 判断不同父节点 同一深度。  关键是如何存放深度，存放父节点
// 递归过程中涉及数据的高度，父节点，可以把高度，父节点放在函数参数中，传给下一层函数
func isCousins(root *TreeNode, x int, y int) bool {
	var xpar, xdep, ypar, ydep int
	var dfs func(node *TreeNode, dep int, x int, y int, par int)
	dfs = func(node *TreeNode, dep int, x int, y int, par int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xpar = par
			xdep = dep
		} else if node.Val == y {
			ypar = par
			ydep = dep
		} else {
			dfs(node.Left, dep+1, x, y, node.Val)
			dfs(node.Right, dep+1, x, y, node.Val)
		}
	}
	dfs(root.Left, 1, x, y, root.Val)
	dfs(root.Right, 1, x, y, root.Val)
	// fmt.Println(xpar,ypar,xdep,ydep)
	return xpar != ypar && xdep == ydep
}

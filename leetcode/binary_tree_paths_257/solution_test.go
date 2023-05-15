package binary_tree_paths_257

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

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	// 初始化path为空 所有的起点
	constructPaths(root, "")
	return paths
}

func constructPaths(node *TreeNode, p string) {
	if node != nil {
		p = p + fmt.Sprint(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, p)
		} else {
			p = p + "->"
			constructPaths(node.Left, p)
			constructPaths(node.Right, p)
		}
	}
}

func s(p *string) {
	*p = *p + "11"
	*p = (*p)[2:]
}

func v(a *[]string) {
	*a = append(*a, "ww")
}

func TestName(t *testing.T) {
	a := ""
	var b *string = &a
	s(b)
	s(b)
	fmt.Println(*b)
	fmt.Println(a)
	fmt.Println(fmt.Sprint(123))

	// 注意 可以添加成功
	*b = a + "11"
	fmt.Println(*b)

	var m = []string{"erw"}
	v(&m)
	fmt.Println(m)

}

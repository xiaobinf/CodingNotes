package flatten_114

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的展开 本质上就是一个前序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var list []*TreeNode
	preOrder(root, &list)

	fmt.Println(list)
	for i := 0; i < len(list)-1; i++ {
		list[i].Left = nil
		list[i].Right = list[i+1]
	}

	//root = list[0]
}

func preOrder(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}

	*list = append(*list, root)
	preOrder(root.Left, list)
	preOrder(root.Right, list)
}

func TestName(t *testing.T) {

}

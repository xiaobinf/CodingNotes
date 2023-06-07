package rightSideView_199

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var vector []*TreeNode
	vector = append(vector, root)

	var result []int
	for len(vector) != 0 {
		result = append(result, vector[len(vector)-1].Val)
		var nextLevel []*TreeNode
		// fmt.Println(len(vector),vector)
		for i := 0; i < len(vector); i++ {
			if vector[i].Left != nil {
				nextLevel = append(nextLevel, vector[i].Left)
			}

			if vector[i].Right != nil {
				nextLevel = append(nextLevel, vector[i].Right)
			}
		}
		vector = nextLevel[:]
	}

	return result
}

func TestName(t *testing.T) {

}

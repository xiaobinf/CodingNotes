package rotate_48

import (
	"fmt"
	"testing"
)

// rotate 转置+左右折叠
func rotate(matrix [][]int) {
	var n = len(matrix)
	var tmp = 0
	// 转置
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	fmt.Println(matrix)
	// 左右折叠
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
	fmt.Println(matrix)
}

func TestName(t *testing.T) {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

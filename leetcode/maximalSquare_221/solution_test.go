package maximalSquare_221

import (
	"fmt"
	"testing"
)

func maximalSquare(matrix [][]byte) int {
	// dp计作正方形的边长
	var dp = make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	dp[0][0] = 0

	var max = 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(min(dp[i-1][j-1], dp[i-1][j]), min(dp[i-1][j-1], dp[i][j-1]))
				if dp[i][j] >= max {
					max = dp[i][j]
				}
			}
		}
	}

	fmt.Println(dp)
	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	maximalSquare([][]byte{{'1', '1', '1', '1'}, {'1', '1', '1', '1'}, {'1', '1', '1', '1'}})
}

package minPathSum_64

import (
	"fmt"
	"testing"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}

	// 初始化dp数组
	var dp = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}

			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}

			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	//fmt.Println(dp)
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func TestName(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 3, 1}}))
	fmt.Println(minPathSum([][]int{{1}, {3}, {1}}))
	fmt.Println(minPathSum([][]int{{1}}))
}

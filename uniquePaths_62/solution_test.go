package uniquePaths_62

import (
	"fmt"
	"testing"
)

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 注意边界条件
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func TestOne(t *testing.T) {
	fmt.Println(uniquePaths(1, 2))
	fmt.Println(uniquePaths(3, 7))
}

package numTrees_96

import (
	"fmt"
	"testing"
)

func numTrees(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	fmt.Println(dp)
	return dp[n]
}

func TestName(t *testing.T) {
	numTrees(5)
	numTrees(3)
}

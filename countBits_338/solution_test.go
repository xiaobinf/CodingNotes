package countBits_338

import (
	"fmt"
	"testing"
)

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}

	return dp
}

func TestName(t *testing.T) {
	fmt.Println(countBits(5))
	fmt.Println(countBits(2))
	fmt.Println(countBits(0))
}

package rob_198

import (
	"fmt"
	"testing"
)

// 这家的最大收益= max（上上家+这家，上家）

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	if len(nums) >= 2 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1]) // 状态转移方程
	}

	fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestOne(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3}))
	fmt.Println(rob([]int{1, 2, 3, 4, 7, 1}))
}

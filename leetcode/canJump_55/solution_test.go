package canJump_55

import (
	"fmt"
	"testing"
)

func canJump(nums []int) bool {
	if nums[0] == 0 && len(nums) > 1 {
		return false
	}

	if nums[0] == 0 {
		return true
	}

	var dp = make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums)-1; i++ {
		dp[i] = max(dp[i-1], nums[i]+i)
		if dp[i] == i {
			return false
		}
	}

	// fmt.Println(dp)
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{0, 2}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

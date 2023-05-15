package lengthOfLIS_300

import (
	"fmt"
	"testing"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	var m = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > m {
			m = dp[i]
		}
	}

	//fmt.Println(dp)
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	//fmt.Println(lengthOfLIS([]int{2}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{2, 3, 6, 4}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

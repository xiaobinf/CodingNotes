package maxProfit_121

import (
	"fmt"
	"testing"
)

// maxProfit 最大股票收益
func maxProfit(prices []int) int {
	// dp公式： dp[i]表示第i天的最大利润，等于昨天的利润+今天与昨天的差价，如果总体小于0，就置为0
	var dp = make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		if dp[i-1]+prices[i]-prices[i-1] >= 0 {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = 0
		}
	}
	//fmt.Println(dp)
	return max(dp)
}

func max(nums []int) int {
	var num = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > num {
			num = nums[i]
		}
	}
	return num
}

func TestName(t *testing.T) {
	fmt.Println(max([]int{3, 2, 1, 9, 3}))
	fmt.Println(maxProfit([]int{7, 1, 3, 5, 4, 6}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

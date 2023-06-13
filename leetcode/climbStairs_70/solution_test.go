package climbStairs_70

import (
	"fmt"
	"testing"
)

func climbStairs1(n int) int {
	// 自顶向下的方式
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := climbStairs(n-1) + climbStairs(n-2)
	return ret
}

//func climbStairs(n int) int {
//	var dp = make([]int, n+2) // 避免溢出
//	dp[1] = 1
//	dp[2] = 2
//
//	for i := 3; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2] // 状态转移方程
//	}
//	fmt.Println(dp)
//	return dp[n]
//}

func TestOne(t *testing.T) {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}

// - - -
// 复习 dp[n] = dp[n-1] + dp[n-2]
func climbStairs(n int) int {
	var dp = make([]int, n+1) // 注意此处 要预先分配数组位置，否则会报数组越界错误
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

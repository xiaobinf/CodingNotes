package numSquares_279

import (
	"fmt"
	"testing"
)

/*这里使用动态规划来做。

一个数=一个数的平方+f(剩余值) 即遍历这个数

定义一个函数f(n)表示我们要求的解。f(n)的求解过程为：
f(n) = 1 + min{
  f(n-1^2), f(n-2^2), f(n-3^2), f(n-4^2), ... , f(n-k^2) //(k为满足k^2<=n的最大的k)
}
*/

func numSquares(n int) int {
	var dp = make([]int, n+1)
	dp[0], dp[1] = 0, 1
	// 自底向上构造
	for i := 1; i <= n; i++ {
		var MIN = 2147483647
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j] < MIN {
				MIN = dp[i-j*j]
			}
		}

		dp[i] = MIN + 1
	}

	return dp[n]
}

func TestName(t *testing.T) {
	fmt.Println(numSquares(2))
}

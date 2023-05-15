package dailyTemperatures_739

import (
	"fmt"
	"testing"
)

//func dailyTemperatures(temperatures []int) []int {
//	var res = make([]int, len(temperatures))
//	for i := 0; i < len(temperatures); i++ {
//		res[i] = 0
//	}
//
//	for i := 0; i < len(temperatures); i++ {
//		for j := i + 1; j < len(temperatures); j++ {
//			if temperatures[j] > temperatures[i] {
//				res[i] = j - i
//				break
//			}
//		}
//	}
//
//	return res
//}

func dailyTemperatures(T []int) []int {
	n := len(T)
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && T[i] >= T[j] {
			if dp[j] > 0 {
				j += dp[j]
			} else {
				j = n
			}
		}
		if j < n {
			dp[i] = j - i
		}
	}
	return dp
}

func TestName(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{3, 2, 5}))
	fmt.Println(dailyTemperatures([]int{2, 3, 5}))
	fmt.Println(dailyTemperatures([]int{4, 3, 3}))
}

package countSubstrings_647

import (
	"fmt"
	"testing"
)

func countSubstrings(s string) int {
	var dp = make([]int, len(s)+1)
	dp[0], dp[1] = 0, 1
	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1]
		for j := 0; j <= i; j++ {
			if judge(s[j:i]) {
				//fmt.Println(s[j:i])
				dp[i] = dp[i] + 1
			}
		}
	}

	//fmt.Println(dp)
	return dp[len(s)]
}

func judge(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func TestName(t *testing.T) {
	fmt.Println(countSubstrings("1"))
	fmt.Println(countSubstrings("123"))
	fmt.Println(countSubstrings("aaa"))
}

package longestPalindrome_5

import (
	"fmt"
	"testing"
)

// longestPalindrome最长回文子串
func longestPalindrome(s string) string {
	var dp = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	var start, end = 0, 0
	var max = 0
	// dp[i][j] = s[i]==s[j]&&dp[i+1][j-1]
	for j := 0; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if i == j {
				continue
			}

			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			}

			if j-i > 1 {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			if j-i >= max && dp[i][j] {
				max = j - i
				start = i
				end = j
			}
		}
	}

	// 求最大回文子串 就是j-1最大 且dp[i][j]=true
	//fmt.Println(dp)
	return s[start : end+1]
}

func TestName(t *testing.T) {
	fmt.Println(longestPalindrome("12121"))
	fmt.Println(longestPalindrome("1"))
}

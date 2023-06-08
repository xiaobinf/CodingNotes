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
		// 初始化 单个字符是回文子串
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

			// 下面两个是核心状态转移方程
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			}

			if j-i > 1 {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			// 保存最大的字串的长度
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

// longestPalindrome 构建二维数组的推导过程
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	// 初始化dp数组
	 for i:=0;i<len(s);i++{
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	 }

	 var max,start,end int
	 // 根据状态转移方程 构建dp数组 只需要构建二位数组一半即可
	 for i:=0;i<len(s);i++ {
		for j:=0;j<i;j++{
			if i==j {
				continue
			}

			if i-j==1&&s[i]==s[j] {
				dp[j][i]=true
			}

			if i-j>1&&dp[j+1][i-1] {
				dp[j][i]=s[i]==s[j]
			}

			if i-j>=max&&dp[j][i]{
				max = i-j
				start = j
				end = i
			}
		}
	 }

	 return s[start: end+1]
}

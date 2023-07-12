package wordBreak_139

import (
	"fmt"
	"testing"
)

// 核心：[0, i - 1] 的字符串可被拆分，当前仅当任一子串 [0, j - 1] 及 [j, i - 1] 可被拆分

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 1 {
		return contains(s, wordDict)
	}

	var dp = make([]bool, len(s))
	for i := 0; i < len(s); i++ { // 判断dp[i]
		for j := 0; j <= i; j++ {
			if j == 0 {
				// 完整的字符串被出现在字典中
				if contains(s[j:i+1], wordDict) {
					fmt.Println(s[j : i+1])
					dp[i] = true
					break
				}
			} else if dp[j-1] && contains(s[j:i+1], wordDict) { // 注意此处下标会溢出 画出详细的表格图
				dp[i] = true
				break
			}
		}
	}

	fmt.Println(dp)
	return dp[len(s)-1]
}

func contains(s string, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		if wordDict[i] == s {
			return true
		}
	}

	return false
}

func TestName(t *testing.T) {
	fmt.Println(wordBreak("ab", []string{"a", "b"}))
	fmt.Println(wordBreak("2311", []string{"1", "23"}))

	fmt.Println(wordBreak("aaa", []string{"aaaa", "aa"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	//fmt.Println(contains("123"[1:3], []string{"1", "23"}))
	fmt.Println("123"[1:1])
}

package is_palindrome_125

import (
	"fmt"
	"regexp"
	"strings"
)

//判断回文字符串 A 65 Z 90 a 97 z 122 '0' 48 '9' 57
//方法1 产生了额外的字符串空间
func isPalindrome1(s string) bool {
	if strings.Trim(s, " ") == "" {
		return true
	}
	// 移除非数字 字母的字符 转小写
	r := regexp.MustCompile(`[^a-zA-Z\d]`)
	s = strings.ToLower(r.ReplaceAllString(s, ""))
	fmt.Println(s)
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isValidChar(c uint8) bool {
	return (65 <= c && c <= 90) || (97 <= c && c <= 122) || (48 <= c && c <= 57)
}

func judgeChar(c1 uint8, c2 uint8) bool {
	if 97 <= c1 && c1 <= 122 {
		c1 = c1 - 32
	}
	if 97 <= c2 && c2 <= 122 {
		c2 = c2 - 32
	}
	return c1 == c2
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	i := 0
	j := len(s) - 1
	for i <= j {
		for i <= j && !isValidChar(s[i]) {
			i++
		}
		for i <= j && !isValidChar(s[j]) {
			j--
		}
		if i <= j && !judgeChar(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

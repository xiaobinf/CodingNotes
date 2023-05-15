package valid_palindrome_680

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
验证回文字串
删除一个字符 判断是否是回文串
*/

// 传入待判定的回文串
func isValid(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			// 对字符串来说只有一次机会来判定， 下面的分支语句可以简化
			if isValid(s[start+1:end+1]) || isValid(s[start:end]) {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func TestName(t *testing.T) {
	assert.Equal(t, isValid("ass"), false)
	assert.Equal(t, isValid("sas"), true)
	assert.Equal(t, isValid("sass"), false)
	assert.Equal(t, isValid(""), true)
	assert.Equal(t, isValid("ssass"), true)
	assert.Equal(t, validPalindrome("ssass"), true)
	assert.Equal(t, validPalindrome("sass"), true)
	assert.Equal(t, validPalindrome("sbabs"), true)
	assert.Equal(t, validPalindrome("sbbbs"), true)
	assert.Equal(t, validPalindrome("ssabs"), false)
	assert.Equal(t, validPalindrome(""), true)
	assert.Equal(t, validPalindrome("abca"), true)
	assert.Equal(t, validPalindrome("abc"), false)
}

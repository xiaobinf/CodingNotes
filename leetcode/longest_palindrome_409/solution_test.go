package longest_palindrome_409

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
最长回文字串
待优化点：可以用用一个128定长数组存放遍历的结果，数组的下标对应字母的ascii码，值对应字母出现的次数。
*/
func longestPalindrome(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	m := map[int32]int{}
	var length int = 0
	// 构造映射存放 字符串元素和计数
	for _, v := range s {
		count, ok := m[v]
		if ok {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}
	// 遍历map
	var flag = true
	for _, v := range m {
		if v%2 == 0 {
			length = length + v
		} else {
			if !flag {
				length = length + v - 1
			} else {
				length = length + v
				flag = false
			}

		}
	}
	return length
}

func TestName(t *testing.T) {
	assert.Equal(t, longestPalindrome("aacccd"), 5)
	assert.Equal(t, longestPalindrome(""), 0)
	assert.Equal(t, longestPalindrome("w"), 1)
	assert.Equal(t, longestPalindrome("ww"), 2)
	assert.Equal(t, longestPalindrome("we"), 1)
	assert.Equal(t, longestPalindrome("wew"), 3)
	assert.Equal(t, longestPalindrome("www"), 3)
	assert.Equal(t, longestPalindrome("YYYYGGGJHHJ"), 11)
}

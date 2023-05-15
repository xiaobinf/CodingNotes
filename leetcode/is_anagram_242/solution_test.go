package is_anagram_242

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 遍历第一个字符串 构建map uint8需要考虑字符串超长的场景
	m := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if !ok {
			m[s[i]] = 1
		} else {
			v++
			m[s[i]] = v
		}
	}

	// 遍历第二个字符串 检查map
	for i := 0; i < len(s); i++ {
		v, ok := m[t[i]]
		if !ok {
			return false
		} else {
			if v == 0 {
				return false
			}
			v--
			m[t[i]] = v
		}
	}

	// 检查map中的元素 应该全部等于0 由于字符串长度相等 这里是多余的。
	// 1。要么多了不存在的字符串 2。要么计数v-1<0
	//for _, v := range m {
	//	if v != 0 {
	//		return false
	//	}
	//}
	return true
}

func Test(t *testing.T) {
	assert.Equal(t, isAnagram("badc", "badc"), true)
	assert.Equal(t, isAnagram("12w", "21w"), true)
	assert.Equal(t, isAnagram("paper", "papre"), true)
	assert.Equal(t, isAnagram("paped", "papeq"), false)
	assert.Equal(t, isAnagram("", "t"), false)
	assert.Equal(t, isAnagram("e", "t"), false)
	assert.Equal(t, isAnagram("121", "211"), true)
	assert.Equal(t, isAnagram("paper", "papre"), true)
}

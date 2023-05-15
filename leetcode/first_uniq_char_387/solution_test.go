package first_uniq_char_387

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	m := make([]int, 26)
	//for i := 0; i < len(m); i++ {
	//	fmt.Println(m[i])
	//}
	for i := 0; i < len(s); i++ {
		m[s[i]-97]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-97] == 1 {
			return i
		}
	}
	return -1
}

func TestName(t *testing.T) {
	assert.Equal(t, firstUniqChar("aabbscscz"), 8)
	assert.Equal(t, firstUniqChar("aabbscsczz"), -1)
	assert.Equal(t, firstUniqChar("a"), 0)
	assert.Equal(t, firstUniqChar("aa"), -1)
	assert.Equal(t, firstUniqChar("aab"), 2)
	assert.Equal(t, firstUniqChar("cabc"), 1)
	assert.Equal(t, firstUniqChar("a"), 0)
}

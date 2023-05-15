package di_string_match_942

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 读懂原理  理解题目
func diStringMatch(s string) []int {
	var arr []int
	l, m := 0, len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			arr = append(arr, l)
			l++
		}
		if s[i] == 'D' {
			arr = append(arr, m)
			m--
		}
		if i == len(s)-1 {
			arr = append(arr, m)
		}
	}
	return arr
}

func TestName(t *testing.T) {
	assert.Equal(t, diStringMatch("III"), []int{0, 1, 2, 3})
	assert.Equal(t, diStringMatch("IDID"), []int{0, 4, 1, 3, 2})
	assert.Equal(t, diStringMatch("DDI"), []int{3, 2, 0, 1})
}

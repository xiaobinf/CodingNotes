package is_isomorphic_205

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := map[uint8]uint8{}
	m2 := map[uint8]uint8{}
	for i := 0; i < len(s); i++ {
		v1, ok1 := m1[s[i]]
		v2, ok2 := m2[t[i]]
		if !ok1 {
			m1[s[i]] = t[i]
		} else if v1 != t[i] {
			return false
		}

		if !ok2 {
			m2[t[i]] = s[i]
		} else if v2 != s[i] {
			return false
		}
	}

	return true
}

func TestSomething(t *testing.T) {
	assert.Equal(t, isIsomorphic("badc", "baba"), false)
	assert.Equal(t, isIsomorphic("12w", "34s"), true)
	assert.Equal(t, isIsomorphic("paper", "title"), true)
	assert.Equal(t, isIsomorphic("paped", "titlt"), false)
	assert.Equal(t, isIsomorphic("", "titlt"), false)
}

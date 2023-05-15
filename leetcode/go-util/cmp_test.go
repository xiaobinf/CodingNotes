package go_util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(INCMP)
	fmt.Println(LT)
	assert.Equal(t, Compare(1, 2), LT)
	assert.Equal(t, Compare(1, 1), EQ)
	assert.Equal(t, Compare(3, 2), GT)
	assert.Equal(t, Compare(3.3, 2), INCMP)
	assert.Equal(t, Compare(3.3, 3.3), EQ)
	assert.Equal(t, Compare(3.3, 3.4), LT)
	assert.Equal(t, Compare([]int{1, 2}, []int{1, 2}), INCMP)
	assert.Equal(t, Compare("asd", "asd"), EQ)

	assert.Equal(t, CompareLE(1, 2), true)
	assert.Equal(t, CompareGE(2, 1), true)
	assert.Equal(t, CompareGE(2, 3), false)
	assert.Equal(t, CompareGT(4, 3), true)
	assert.Equal(t, CompareLT(1, 3), true)
}

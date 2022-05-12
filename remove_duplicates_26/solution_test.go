package remove_duplicates_26

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	start, end := 1, 1
	for end < len(nums) {
		if nums[end] != nums[end-1] {
			nums[start] = nums[end]
			start++
			end++
		} else {
			end++
		}
	}
	return start
}

func TestName(t *testing.T) {
	assert.Equal(t, 1, 1)
}

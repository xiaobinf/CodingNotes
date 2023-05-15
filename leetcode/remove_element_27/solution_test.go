package remove_element_27

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func removeElement(nums []int, val int) int {
  start, end := 0, len(nums)
  for start < end {
    if nums[start] == val {
      nums[start] = nums[end-1]
      end--
    } else {
      start++
    }
  }
  return start
}

func TestName(t *testing.T) {
  assert.Equal(t, 1, 1)
}

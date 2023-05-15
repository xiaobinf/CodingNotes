package maxArea_11

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	res := 0
	for left < right {
		h := min(height[left], height[right])
		area := h * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 3, 2, 5, 25, 24, 5}))
}

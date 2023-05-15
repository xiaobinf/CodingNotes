package longestConsecutive_128

import (
	"fmt"
	"testing"
)

func longestConsecutive(nums []int) int {
	var m = make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	maxLen := 0
	for num := range m {
		if !m[num-1] {
			start := num
			mlen := 1
			for m[start+1] {
				start++
				mlen++
			}
			maxLen = max(maxLen, mlen)
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestName(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{100}))
}

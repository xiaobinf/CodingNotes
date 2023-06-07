package findKthLargest_215

import (
	"fmt"
	"testing"
)

// findKthLargest 第k大的元素 快速排序的分治法
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	l := 0
	r := len(nums) - 1
	var index int
	for {
		index = partition(nums, l, r)
		if index == target {
			return nums[target]
		} else if index < target {
			l = index + 1
		} else {
			r = index - 1
		}
	}
}

// 划分 找到pivot的一个排序后固定的位置
func partition(nums []int, l, r int) int {
	pivot := nums[l]
	for l < r {
		for l < r && pivot <= nums[r] {
			r--
		}
		nums[l] = nums[r]

		for l < r && pivot >= nums[l] {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}

func TestName(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 4, 2}, 2))
}

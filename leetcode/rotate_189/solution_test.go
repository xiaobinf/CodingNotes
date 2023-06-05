package rotate_189

import (
	"fmt"
	"testing"
)

func rotate(nums []int, k int) {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[(i+k)%len(nums)] = nums[i]
	}

	copy(nums, nums1)
}

func TestName(t *testing.T) {
	var nums = []int{1, 2, 3, 4}
	rotate(nums, 2001)
	fmt.Println(nums)
}

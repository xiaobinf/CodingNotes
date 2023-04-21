package sortColors_75

import (
	"fmt"
	"testing"
)

func sortColors(nums []int) {
	var c0, c1, c2 = 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c0++
		} else if nums[i] == 1 {
			c1++
		} else {
			c2++
		}
	}

	for i := 0; i < c0; i++ {
		nums[i] = 0
	}

	for i := c0; i < c0+c1; i++ {
		nums[i] = 1
	}

	for i := c0 + c1; i < len(nums); i++ {
		nums[i] = 2
	}
}

func TestName(t *testing.T) {
	var arr = []int{2, 1, 0, 0, 2, 1, 0, 1}
	sortColors(arr)
	fmt.Println(arr)
}

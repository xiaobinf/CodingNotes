package findDuplicate_287

import (
	"fmt"
	"testing"
)

// 错误解法
func findDuplicate(nums []int) int {
	var num = (1 + len(nums) - 1) * (len(nums) - 1) / 2
	var num1 = 0
	for i := 0; i < len(nums); i++ {
		num1 = num1 + nums[i]
	}

	return num1 - num
}

func findDuplicate1(nums []int) int {
	var m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			return nums[i]
		} else {
			m[nums[i]]++
		}
	}

	return -1
}

func TestName(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 2, 3, 3}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 2}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(findDuplicate([]int{1, 1}))

	fmt.Println(findDuplicate1([]int{1, 2, 3, 3}))
	fmt.Println(findDuplicate1([]int{1, 2, 3, 2}))
	fmt.Println(findDuplicate1([]int{1, 2, 3, 1}))
	fmt.Println(findDuplicate1([]int{1, 1}))
}

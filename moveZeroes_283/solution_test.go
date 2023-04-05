package moveZeroes_283

import (
	"testing"
)

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
	//fmt.Println(nums)
}

func TestName(t *testing.T) {
	moveZeroes([]int{1, 0, 4, 5, 0})
	moveZeroes([]int{0, 0, 4, 5, 1})
	moveZeroes([]int{0, 0, 4, 5, 0})
	moveZeroes([]int{0, 0, 4, 5, 3})
	moveZeroes([]int{0, 0})
}

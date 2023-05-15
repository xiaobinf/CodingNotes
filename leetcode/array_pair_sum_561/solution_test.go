package array_pair_sum_561

import (
	"fmt"
	"sort"
	"testing"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}

func TestName(t *testing.T) {
	fmt.Println(arrayPairSum([]int{1, 3, 5, 4, 2, 9}))
}

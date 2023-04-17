package findDisappearedNumbers_448

import (
	"fmt"
	"testing"
)

// 超时
//func findDisappearedNumbers(nums []int) []int {
//	var res []int
//	for i := 0; i < len(nums); i++ {
//		var j int
//		for j = 0; j < len(nums); j++ {
//			if i+1 == nums[j] {
//				break
//			}
//		}
//
//		if j == len(nums) {
//			res = append(res, i+1)
//		}
//
//	}
//	return res
//}

func findDisappearedNumbers(nums []int) []int {
	var res = make([]int, len(nums))
	var ret []int
	for i := 0; i < len(res); i++ {
		res[i] = 0
	}

	for i := 0; i < len(res); i++ {
		res[nums[i]-1] = 1
	}

	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}

func TestName(t *testing.T) {
	fmt.Println(findDisappearedNumbers([]int{3, 2, 4, 5, 5}))
	fmt.Println(findDisappearedNumbers([]int{3, 1, 4, 5, 5}))
	fmt.Println(findDisappearedNumbers([]int{3, 4, 4, 3, 5}))
}

package subsets_78

import (
	"fmt"
	"testing"
)

//func subsets(nums []int) [][]int {
//	var res [][]int
//	res = append(res, []int{})
//	for i := 0; i < len(nums); i++ {
//		for _, subset := range res {
//			res = append(res, append([]int{nums[i]}, subset...))
//		}
//	}
//
//	fmt.Println(res)
//	return res
//}

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		var length = len(res)
		for j := 0; j < length; j++ {
			//fmt.Println(res[j], nums[i])
			tmp := append([]int{nums[i]}, res[j]...)
			res = append(res, tmp)
		}
	}

	fmt.Println(res)
	return res
}

func TestName(t *testing.T) {
	subsets([]int{9, 0, 3, 5, 7})
}

package permute_46

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack(nums, &res, track)
	return res
}

func backtrack(nums []int, res *[][]int, track []int) {
	if len(track) == len(nums) {
		// 需要重新构造切片对象
		*res = append(*res, append([]int{}, track...))
		track = []int{}
		return
	}

	for i := 0; i < len(nums); i++ {
		// 判断是否存在重复
		if contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, res, track)
		track = track[:len(track)-1]
	}
}

func contains(track []int, tmp int) bool {
	for i := 0; i < len(track); i++ {
		if track[i] == tmp {
			return true
		}
	}
	return false
}

func TestOne(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}

package combinationSum_39

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	sort.Ints(candidates)
	backtrack(candidates, target, &res, &track)
	fmt.Println("res :::", res)
	return res
}

func backtrack(candidates []int, target int, res *[][]int, track *[]int) {
	// 递归退出条件
	if sum(*track) == target {
		*res = append(*res, append([]int{}, *track...))
		return
	}

	if sum(*track) > target {
		return
	}

	// 遍历子节点
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue // 跳过相同的元素[1, 2, 2] 会出现重复的组合
		}
		*track = append(*track, candidates[i])
		fmt.Println(candidates[i:])
		backtrack(candidates[i:], target, res, track) // 从当前元素开始搜索 防止后面又重复
		*track = (*track)[:len(*track)-1]             // 删除当前元素
	}
}

func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func sum(num []int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	return sum
}

func modifySlice(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

func TestName(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", slice)
	modifySlice(&slice)
	fmt.Println("After:", slice)

	combinationSum([]int{1, 2, 2, 3, 4, 5}, 5)

}

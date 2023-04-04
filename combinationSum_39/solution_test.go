package combinationSum_39

import (
	"fmt"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	backtrack(candidates, target, &res, track)
	fmt.Println("res end", res)
	return res
}

func backtrack(candidates []int, target int, res *[][]int, track []int) {
	if len(track) <= len(candidates) && sum(track) == target {
		fmt.Println("track", track)
		*res = append(*res, track)
		fmt.Println("res", res)
		track = []int{}
		return
	}

	if len(track) == len(candidates) && sum(track) != target {
		track = []int{}
		return
	}

	for i := 0; i < len(candidates); i++ {
		// todo check contains
		//fmt.Println(candidates[i])
		track = append(track, candidates[i])
		backtrack(candidates, target, res, track)
		//fmt.Println(track)
		track = (track)[:len(track)-1]
	}
}

func sum(num []int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	return sum
}

func TestName(t *testing.T) {
	combinationSum([]int{1, 2}, 3)
}

package backtracking

import (
	"fmt"
	"testing"
)

/**
在全排列问题中，我们可以通过交换数组中的元素来生成不同的排列。具体来说，我们可以从数组的第一个位置开始，依次将每个元素与当前位置交换，并递归生成剩余元素的全排列。在递归过程中，我们需要记录已经生成的排列，并在搜索到数组的末尾时将其添加到结果中。

在回溯算法中，我们需要注意以下几点：

递归终止条件：当搜索到数组的末尾时，将当前排列添加到结果中，并返回上一层递归。

递归过程中的剪枝：在交换元素之前，我们需要判断当前元素是否已经在排列中出现过。如果已经出现过，则不需要再次交换，以避免生成重复的排列。

回溯过程中的状态重置：在回溯到上一层递归时，我们需要将数组中的元素交换回来，以便继续搜索其他可能的解。
*/
// permute 返回给定数组的全排列
func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, &res, 0)
	return res
}

func backtrack(nums []int, res *[][]int, start int) {
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrack(nums, res, start+1)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

/**
给出另一种思路 交换逻辑比较复杂 给出简单方式
理解为构建了一个完全二叉树，自根向下逐渐遍历，track记录路径。 回退时删除这个点。
*/

func permute1(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack1(nums, &res, track)
	return res
}

func backtrack1(nums []int, res *[][]int, track []int) {
	if len(track) == len(nums) {
		*res = append(*res, append([]int{}, track...))
		track = []int{}
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack1(nums, res, track)
		track = track[:len(track)-1]
	}

}

func contains(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}

func TestName(t *testing.T) {
	fmt.Println(permute1([]int{1, 2, 3, 4, 5}))
}

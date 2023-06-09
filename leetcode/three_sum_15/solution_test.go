package three_sum_15

import (
	"sort"
	"testing"
)

// 一次遍历+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	sort.Ints(nums)
	// 三个指针 i j k
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 每一次循环 k都是从尾部开始
		k := n - 1
		target := -1 * nums[i]
		for j := i + 1; j < n; j++ {
			// 去除重复解
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return result
}

func TestName(t *testing.T) {
	//twoSum([]int{-1, 1, 2, 4, 4, 5, 4, 5, 7, 8, 10, 11}, 9)
	//twoSum([]int{-1, -1, 0, 1, 2}, 4)
	//threeSum([]int{-1, 1, 0, -2, -2, -4, -4, -4, 4, 5, 4, -5, -5, -5, 7, -7, -8, -10, 8, 10, -1, 11})
	threeSum([]int{-4, -1, -1, 0, 1, 2})
	threeSum([]int{-4, -4, -4, -1, -1, -2, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2})
	threeSum([]int{})
	threeSum([]int{0})
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	// 定义返回值
	var res = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		var target = -1 * nums[i]
		k := len(nums) - 1
		//固定住i，对j，k做移动
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 固定住j 对k做移动
			for k > j && nums[j]+nums[k] > target {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target {
				res = append(res, []int{i, j, k})
			}

		}
	}

	return res
}

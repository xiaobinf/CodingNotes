package main

import (
	"fmt"
)

//func searchRange(nums []int, target int) []int {
//	start, end := -1, -1
//	n := len(nums)
//	left, right := 0, n-1
//	var mid int
//	for left <= right {
//		mid = (left + right) / 2
//		if nums[mid] == target {
//			// 左右扩展
//			start, end = mid, mid
//			// 注意边界 -1可能越界
//			for start > 0 && nums[start-1] == target {
//				start--
//			}
//
//			for end < n-1 && nums[end+1] == target {
//				end++
//			}
//
//			break
//		} else if nums[mid] < target {
//			left++
//		} else {
//			right--
//		}
//	}
//
//	return []int{start, end}
//}

func main() {
	fmt.Println(123)
}

// 复习

func searchRange(nums []int, target int) []int {
	var left, right = 0, len(nums) - 1
	var start, end = -1, -1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		start, end = mid, mid
		if nums[mid] == target {
			// 左右跳转
			for start > 0 && nums[start] == target { // 这里越过界了  返回时需要调整
				start--
			}

			for end < len(nums)-1 && nums[end] == target {
				end++
			}
		} else if nums[mid] < target {
			left++
		} else {
			right--
		}
	}

	return []int{start + 1, end - 1}
}

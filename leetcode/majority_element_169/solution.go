package main

import "fmt"

// 考虑两两抵消
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var count = 1
	//num := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			count--
			if count == 0 {
				nums[0] = nums[i+1]
				i++
				count++
			} else {

			}
		} else {
			count++
		}
	}
	return nums[0]
}

func main() {
	nums1 := []int{33, 4, 33, 4, 4, 4, 4}
	fmt.Println(majorityElement(nums1))
}

// 多数元素  可以使用抵消的算法 这样可以一次遍历实现 nums[0]存放当前最多的元素
func majorityElement1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			count--
			if count == 0 {
				nums[0] = nums[i+1]
				i++ // 需要向后移动一个位置
				count++
			}
		} else {
			count++
		}
	}

	return nums[0]
}

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

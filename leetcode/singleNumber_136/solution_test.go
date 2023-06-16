package singleNumber_136

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i] // 异或操作
	}

	return res
}

func TestOne(t *testing.T) {
	fmt.Println(singleNumber([]int{1}))
	fmt.Println(singleNumber([]int{3, 2, 2}))
}


// 只出现一次的数字 暴力解法
func singleNumber(nums []int) int {
	var m =make(map[int]int)
	for i:=0;i<len(nums);i++{
		m[nums[i]]++
	}

	for k,v := range m {
		if v==1{
			return k
		}
	}

	return 0
}
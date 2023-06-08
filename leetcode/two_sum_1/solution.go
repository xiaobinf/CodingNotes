package main

import "fmt"

/*
本质上通过哈希表记住 值和下标索引。
当检查到满足条件的值，就返回对应的下标
*/
// twoSum1 暴力解法
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hashmap
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{i, p}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	fmt.Println(twoSum2(nums1, 5))
}


// 练习
func twoSum(nums []int, target int) []int {
	var h = make(map[int]int)
	for i:=0;i<len(nums);i++{
		if value,ok:=h[target-nums[i]]; ok{
			return []int{i, value}
		}
		h[nums[i]] = i
	}
	
	return []int{}
}
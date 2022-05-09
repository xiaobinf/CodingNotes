package main

import "fmt"

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

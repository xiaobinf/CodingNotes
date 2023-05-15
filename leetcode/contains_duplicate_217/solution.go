package main

import "fmt"

// 判断在map中是否存在
func inMap(target string, _map map[string]int) bool {
	_, ok := _map[target]
	if ok {
		return true
	}
	return false
}

// hash判断即可
func containsDuplicate(nums []int) bool {
	l := len(nums)
	if l == 0 || l == 1 {
		return false
	}
	var m = make(map[int]struct{})
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(containsDuplicate(nums))
}

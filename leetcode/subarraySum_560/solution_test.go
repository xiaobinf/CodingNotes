package subarraySum_560

import "testing"

// subarraySum 子数组和为k
func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

func TestName(t *testing.T) {

}

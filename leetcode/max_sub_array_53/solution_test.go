package max_sub_array_53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 算法本质，当一个子序列逐个相加和小于0，说明该序列对后续的序列和的贡献为0，需要重新定义子序列的头部
// 以序列和小于0为分界点，一次循环，处理完所有的的子问题。
// 注意数组元素可能全部是负数,所以maxSum初始化为第一个元素的值
func maxSubArray(nums []int) int {
	var maxSum, tmpSum = nums[0], 0
	for i := 0; i < len(nums); i++ {
		// 每一步都是计算了当前的子序列的和
		if tmpSum > 0 {
			tmpSum = tmpSum + nums[i]
		} else {
			// 从头开始计算
			tmpSum = nums[i]
		}
		// 每一步保存最大和
		maxSum = maxInt(maxSum, tmpSum)
	}
	return maxSum
}

func TestName(t *testing.T) {
	var arr0 = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var arr1 = []int{1}
	var arr2 = []int{5, 4, -1, 7, 8}
	var arr3 = []int{-5, -4, -1, -7, -8}
	var arr4 = []int{-3, -2, 0, -1}
	assert.Equal(t, maxSubArray(arr0), 6)
	assert.Equal(t, maxSubArray(arr1), 1)
	assert.Equal(t, maxSubArray(arr2), 23)
	assert.Equal(t, maxSubArray(arr3), -1)
	assert.Equal(t, maxSubArray(arr4), 0)

}

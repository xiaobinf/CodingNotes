package trap_42

import (
	"fmt"
	"testing"
)

// trap 接雨水 按列计算
func trap(height []int) int {
	var sum int
	for i := 1; i < len(height)-1; i++ {
		var maxLeft, maxRight int
		for j := 0; j < i; j++ {
			maxLeft = max(maxLeft, height[j])
		}

		for k := i + 1; k < len(height); k++ {
			maxRight = max(maxRight, height[k])
		}
		var minLR = min(maxLeft, maxRight)
		if minLR <= height[i] {
			continue
		} else {
			sum = sum + minLR - height[i]
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestName(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

// 复习 核心按列计算 每一列等于左右最低的值
func trap1(height []int) int {
	var sum int
	for i := 1; i < len(height)-1; i++ {
		var maxLeft, maxRight int //局部定义
		// 获取左侧的最大值
		for j := 0; j < i; j++ {
			maxLeft = max(height[j], maxLeft)
		}

		// 获取右侧的最大值
		for j := i + 1; j < len(height); j++ {
			maxRight = max(height[j], maxRight)
		}

		// 计算单列面积
		minLR := min(maxLeft, maxRight)
		// 放不住水
		if minLR <= height[i] {
			continue
		} else {
			sum = sum + minLR - height[i]
		}
	}

	return sum
}

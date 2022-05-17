package largest_perimeter_976

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/* 本质上从数组中找3个数组成三角形，且这三个数和最大 */
/*
为什么排序遍历相邻元素可行，有没有可能最优解为非相邻元素？（不会）

证明：反证 假设 a , b, c 为最优解，且存在a',b',满足 a < a' < b < b' < c（存在非相邻元素）

    由于 a + b > c，a < a', 有 a' + b > c，(a', b, c)优于(a, b, c),与(a, b, c)为最优解矛盾，故不存在a'
    b'同理不存在 由于 a + b > c, b < b'，有a + b' > c，(a, b, c)为最优解矛盾，故不存在b'

因此最优解一定为排序后相邻元素
所以只需要找到判断连续的三个数字即可
*/
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func judge(arr0, arr1, arr2 int) bool {
	if arr1+arr2 > arr0 {
		return true
	}
	return false
}

func largestPerimeter_1(nums []int) int {
	sort.Ints(nums)
	// 找三个较大的数
	for i := len(nums) - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			if nums[j]*2 <= nums[i] {
				break
			}
			for k := j - 1; k >= 0; k-- {
				if judge(nums[i], nums[j], nums[k]) {
					//fmt.Println(nums[i], " ", nums[j], " ", nums[k])
					return nums[i] + nums[j] + nums[k]
				}
				break
			}
		}
	}
	return 0
}

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	// 找三个较大的数
	for i := len(nums) - 1; i >= 2; i-- {
		if judge(nums[i], nums[i-1], nums[i-2]) {
			//fmt.Println(nums[i], " ", nums[j], " ", nums[k])
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func TestName(t *testing.T) {
	//assert.Equal(t, judge(1, 2, 3), false)
	//assert.Equal(t, judge(3, 3, 3), true)
	//assert.Equal(t, judge(1, 3, 3), true)
	//assert.Equal(t, judge(5, 2, 3), false)
	assert.Equal(t, largestPerimeter([]int{5, 4, 3}), 12)
	assert.Equal(t, largestPerimeter([]int{5, 4, 1}), 0)
	assert.Equal(t, largestPerimeter([]int{5, 4, 5, 6, 7}), 18)
	assert.Equal(t, largestPerimeter([]int{5, 4, 2, 7}), 16)
	assert.Equal(t, largestPerimeter([]int{3, 6, 2, 3}), 8)
	assert.Equal(t, largestPerimeter([]int{1, 1, 1}), 3)
	assert.Equal(t, largestPerimeter([]int{1, 2, 1, 1, 1, 1, 1}), 3)
}

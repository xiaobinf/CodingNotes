package balanced_string_split_1221

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func isBalancedString(s string) bool {
//	num1, num2 := 0, 0
//	for _, v := range s {
//		if v == 76 {
//			num1++
//		} else {
//			num2++
//		}
//	}
//	return num1 == num2
//}

//func balancedStringSplit(s string) int {
//	count := 0
//	for i := 0; i < len(s)-1; i++ {
//		for j := i + 1; j < len(s); j++ {
//			fmt.Println(count)
//			if isBalancedString(s[i : j+1]) {
//				count++
//				// 此处不能直接加1，会出现重复
//				i = j
//				j = i
//				break
//			}
//		}
//	}
//	return count
//}

/*
维护差量
*/
func balancedStringSplit(s string) (ans int) {
	d := 0
	for _, ch := range s {
		if ch == 'L' {
			d++
		} else {
			d--
		}
		if d == 0 {
			ans++
		}
	}
	return
}

func TestName(t *testing.T) {
	//assert.Equal(t, isBalancedString("RLRLRL"), true)
	//assert.Equal(t, isBalancedString("RRRLLL"), true)
	//assert.Equal(t, isBalancedString("RL"), true)
	//assert.Equal(t, isBalancedString("RLL"), false)
	//assert.Equal(t, isBalancedString("R"), false)
	//assert.Equal(t, isBalancedString("RRLLLR"), true)
	assert.Equal(t, balancedStringSplit("LR"), 1)
	assert.Equal(t, balancedStringSplit("LRLR"), 2)
	assert.Equal(t, balancedStringSplit("LLRR"), 1)
	assert.Equal(t, balancedStringSplit("RLRLRRLL"), 3)
	assert.Equal(t, balancedStringSplit("LLLLRRRR"), 1)
	assert.Equal(t, balancedStringSplit("RRRLLLRLRL"), 3)
	assert.Equal(t, balancedStringSplit("RLRRRLLRLL"), 2)
	assert.Equal(t, balancedStringSplit("RLRRLLLRLLRR"), 4)

}

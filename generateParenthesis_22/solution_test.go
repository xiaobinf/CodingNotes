package generateParenthesis_22

import (
	"fmt"
	"testing"
)

var bracket = "()"

// generateParenthesis 构造六层括号的解空间树  生成全排列
func generateParenthesis(n int) []string {
	var res, ret []string
	var track string
	backtrack(n, &res, track)
	// 过滤不满足条件的组合
	for i := 0; i < len(res); i++ {
		if judgeClose(res[i]) {
			ret = append(ret, res[i])
		}
	}
	return ret
}

func backtrack(n int, res *[]string, track string) {
	if len(track) == 2*n {
		// 生成一个解后的终止逻辑
		*res = append(*res, track)
		track = ""
		return
	}
	for i := 0; i < 2; i++ {
		// 剪枝逻辑 此处没有
		track = track + string(bracket[i])
		backtrack(n, res, track)
		track = track[:len(track)-1]
	}
}

// judgeClose 判断括号是否闭合 左括号入栈  右括号弹出判断
func judgeClose(res string) bool {
	stack := make([]rune, 0, len(res))
	for i := 0; i < len(res); i++ {
		if res[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func TestName(t *testing.T) {
	fmt.Println(generateParenthesis(2))
	fmt.Println(judgeClose("()"))
	fmt.Println(judgeClose("(("))
	fmt.Println(judgeClose("()()"))
	fmt.Println(judgeClose("())("))
}

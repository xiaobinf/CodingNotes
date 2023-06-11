package generateParenthesis_22

import (
	"fmt"
	"testing"
)

//var bracket = "()"

// generateParenthesis 构造六层括号的解空间树  生成全排列
//func generateParenthesis(n int) []string {
//	var res, ret []string
//	var track string
//	backtrack(n, &res, track)
//	// 过滤不满足条件的组合
//	for i := 0; i < len(res); i++ {
//		if judgeClose(res[i]) {
//			ret = append(ret, res[i])
//		}
//	}
//	return ret
//}
//
//func backtrack(n int, res *[]string, track string) {
//	if len(track) == 2*n {
//		// 生成一个解后的终止逻辑
//		*res = append(*res, track)
//		track = ""
//		return
//	}
//	for i := 0; i < 2; i++ {
//		// 剪枝逻辑 此处没有
//		track = track + string(bracket[i])
//		backtrack(n, res, track)
//		track = track[:len(track)-1]
//	}
//}

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
	//fmt.Println(generateParenthesis(2))
	fmt.Println(judgeClose("()"))
	fmt.Println(judgeClose("(("))
	fmt.Println(judgeClose("()()"))
	fmt.Println(judgeClose("())("))
	fmt.Println("-----------------")

	fmt.Println(judge("()"))
	//fmt.Println(judge("(("))
	//fmt.Println(judge("()()"))
	//fmt.Println(judge("())("))
	fmt.Println("-----------------")

	fmt.Println("()"[0] == uint8('('))
	fmt.Println("()"[0] == '(')
	fmt.Println("()"[1])
}

// 复习  括号生成  使用回溯算法 可以在添加解集合的时候判断合法性

// 判断函数
func judge(s string) bool {
	//var stack = make([]uint8, len(s)) // 此处出错 预先给stack赋值
	var stack []uint8
	if s[0] == uint8(')') {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == uint8('(') {
			// 进栈
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				fmt.Println(stack)
				stack = stack[:len(stack)-1]
				fmt.Println(stack)

			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		fmt.Println(13)
		return false
	}
}

func generateParenthesis(n int) []string {
	var res []string
	var track string
	backtrack(&res, track, n)
	return res
}

func backtrack(res *[]string, track string, n int) {
	if len(track) == 2*n {
		// 终止逻辑
		if judge(track) {
			*res = append(*res, track)
			track = ""
			return
		} else {
			track = ""
			return
		}

	}

	s := "()"
	for i := 0; i < len(s); i++ {
		// 无剪枝逻辑
		track = track + string(s[i])
		backtrack(res, track, n)
		track = track[:len(track)-1]
	}
}

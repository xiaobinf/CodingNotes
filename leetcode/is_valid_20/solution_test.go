package is_valid_20

//有效的括号 '('，')'，'{'，'}'，'['，']'
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}
	var stk []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk = append(stk, s[i])
		}
		if s[i] == ')' {
			if len(stk) != 0 && stk[len(stk)-1] == '(' {
				stk = stk[0 : len(stk)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if len(stk) != 0 && stk[len(stk)-1] == '[' {
				stk = stk[0 : len(stk)-1]
			} else {
				return false
			}
		} else if s[i] == '}' {
			if len(stk) != 0 && stk[len(stk)-1] == '{' {
				stk = stk[0 : len(stk)-1]
			} else {
				return false
			}
		}
	}
	if len(stk) == 0 {
		return true
	} else {
		return false
	}
}

// 复习 合法的括号
// 核心思路 1.遇左入栈 遇右出栈 2.使用列表来模拟栈

func isValid(s string) bool {
	var stack []uint8
	if len(s)%2!=0{
		return false
	}

	if len(s)==0{
		return true
	}

	if s[0] == '}' || s[0] == ']' || s[0] == ')' {
		return false
	}

	for i:=0; i< len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		}

		if s[i] == '}' {
			if len(stack)>0 && stack[len(stack)-1]=='{'{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}

		if s[i] == ']' {
			if len(stack)>0 && stack[len(stack)-1]=='['{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}

		if s[i] == ')' {
			if len(stack)>0 && stack[len(stack)-1]=='('{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}
	}

	if len(stack)==0{
		return true
	}else{
		return false
	}
}
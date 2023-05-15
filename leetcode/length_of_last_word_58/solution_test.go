package length_of_last_word_58

// 核心： 从后往前遍历，计算非空字符数 第二次遇到空字符停止
func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length != 0 {
				return length
			} else {
				length = 0
			}
		} else {
			length++
		}
	}
	return length
}

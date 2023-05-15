package plus_one_66

// 可以拓展为加n
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			// 核心点：如果没有进位 可以直接返回
			return digits
		}
		digits[i] = 0
	}
	// 如果跳出for循环 那么全是9
	// 注意数组的构造方式, 这里其实就是初始化 [0,0,0,0,0...0]
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

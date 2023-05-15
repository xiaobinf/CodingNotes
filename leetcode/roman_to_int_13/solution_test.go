package roman_to_int_13

func romanToInt(s string) int {
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var num = 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && m[string(s[i])] < m[string(s[i+1])] {
			num -= m[string(s[i])]
		} else {
			num += m[string(s[i])]
		}
	}
	return num
}

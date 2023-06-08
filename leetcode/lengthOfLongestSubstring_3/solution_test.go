// lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int)
	r, ans := -1,0
	for i:=0;i<len(s);i++ {
		// i向后移动的时候，需要删除原来map中的最左侧的字符
		if i!=0{
			delete(m, s[i-1])
		}
		// 如果我们依次递增地枚举子串的起始位置，那么子串的结束位置也是递增的！即r的位置是递增的
		// r不需要回退的
		for r+1<len(s)&&m[s[r+1]]==0 {
			m[s[r+1]]++
			r++
		}

		ans = max(ans, r-i+1)
	}

    return ans
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

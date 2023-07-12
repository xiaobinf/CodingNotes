func firstUniqChar(s string) byte {
	if s==""{
		return ' '
	}
	var ret byte
	var m = make(map[byte][int])
	for i:=0;i<len(s);i++{
		m[s[i]]+=1
	}

	for i:=0;i<len(s);i++{
		if m[s[i]]==1{
			ret = s[i]
			break
		}
	}

	return ret
}

Line 3: Char 22: empty rune literal or unescaped ' (solution.go)
Line 30: Char 19: undefined: bufio (solution.go)
Line 30: Char 35: undefined: os (solution.go)
Line 31: Char 19: undefined: os (solution.go)
Line 41: Char 35: undefined: io (solution.go)
Line 48: Char 24: undefined: strings (solution.go)
Line 71: Char 17: undefined: fmt (solution.go)
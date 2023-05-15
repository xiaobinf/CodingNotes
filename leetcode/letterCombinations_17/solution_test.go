package letterCombinations_17

import (
	"fmt"
	"testing"
)

//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//输入：digits = "2"
//输出：["a","b","c"]
//输入：digits = ""
//输出：[]

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	// 构造数字 字母映射表
	var table = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	var res []string
	var track string
	// 这个也是全排列 但是要记住下标 因为解空间树的每一层都不一致
	backtrack(digits, &res, track, &table, 0)
	return res
}

func backtrack(digits string, res *[]string, track string, table *map[int]string, index int) {
	if len(track) == len(digits) {
		*res = append(*res, track)
		track = ""
		return
	}

	for i := 0; i < len((*table)[int(digits[index]-'0')]); i++ {
		track = track + string((*table)[int(digits[index]-'0')][i])
		backtrack(digits, res, track, table, index+1)
		track = track[:len(track)-1]
	}
}

func TestName(t *testing.T) {
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("22"))
	fmt.Println("16622133"[2] - '0')
}

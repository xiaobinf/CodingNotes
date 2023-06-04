package findAnagrams_438

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	var arr []int

	for i := 0; i < len(s)-len(p); i++ {
		if judge1(s[i:i+len(p)], p) {
			arr = append(arr, i)
		}
	}

	return arr
}

// judge 判断异位词 该方式会超时
func judge(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s3 := strings.Split(s1, "")
	sort.Strings(s3)

	s4 := strings.Split(s2, "")
	sort.Strings(s4)

	if strings.Join(s3, "") == strings.Join(s4, "") {
		return true
	} else {
		return false
	}
}

// judge 判断异位词1
func judge1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var num1 [26]int
	var num2 [26]int

	for i := 0; i < len(s1); i++ {
		num1[s1[i]-'a']++
	}

	for i := 0; i < len(s2); i++ {
		num2[s2[i]-'a']++
	}

	if num1 == num2 {
		return true
	} else {
		return false
	}
}

func TestName(t *testing.T) {
	fmt.Println(judge("123", "231"))
	fmt.Println(judge("1213", "231"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

}

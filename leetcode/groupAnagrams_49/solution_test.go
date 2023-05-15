package groupAnagrams_49

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)
	var res [][]string

	for i := 0; i < len(strs); i++ {
		s := _sort(strs[i])
		fmt.Println(s)
		m[s] = append(m[s], strs[i])
	}

	for k, _ := range m {
		res = append(res, m[k])
	}

	return res
}

func _sort(s string) string {
	a := strings.Split(s, "")
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	result := strings.Join(a, "")

	return result
}

func TestName(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

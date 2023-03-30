package backtracking

import (
	"fmt"
	"testing"
)

/**
理解为构建了一个完全二叉树，自根向下逐渐遍历，track记录路径。 回退时删除这个点。
这里与常规的全排列有区别：1、字符可以重复，不能以字符的值的维度去重，要以下表维度去重 2、去完重排列后 还需要整体去重
*/

func permutation(s string) []string {
	var res []string
	var track string
	var indexs []int // 记录下标 用来去重
	backtrack(s, &res, track, indexs)
	// 存在重复元素 需要去重
	m := make(map[string]bool)
	for _, str := range res {
		m[str] = true
	}

	// 将 map 中的键转换为字符串切片
	result := make([]string, 0, len(m))
	for str := range m {
		result = append(result, str)
	}

	return result
}

func backtrack(s string, res *[]string, track string, indexs []int) {
	// 定义终止条件
	if len(s) == len(track) {
		*res = append(*res, track)
		track = ""
		indexs = []int{}
		return
	}

	// 判断路径
	for i := 0; i < len(s); i++ {
		// 剪枝逻辑 如果是前一层的这个位置的字符，就应该扔掉
		// 所以需要记住之前存放的位置index
		if contains(indexs, i) {
			continue
		}

		// 添加路径
		track = track + string(s[i])
		indexs = append(indexs, i)
		backtrack(s, res, track, indexs)
		// 回溯
		track = track[:len(track)-1]
		indexs = indexs[:len(indexs)-1]
	}
}

func contains(indexs []int, num int) bool {
	for i := 0; i < len(indexs); i++ {
		if indexs[i] == num {
			return true
		}
	}
	return false
}

func TestName(t *testing.T) {
	fmt.Println(permutation("aba"))
}

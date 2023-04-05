package go_util

func DeleteSliceElms(sl []int, elms ...int) []int {
	// 先将元素转为 set。
	m := make(map[int]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 过滤掉指定元素。
	res := make([]int, 0, len(sl))
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

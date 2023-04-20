package merge_56

import (
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for _, interval := range intervals {
		if len(res) == 0 || interval[0] > res[len(res)-1][1] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(interval[1], res[len(res)-1][1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestName(t *testing.T) {

}

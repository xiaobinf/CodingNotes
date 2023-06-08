package setZeroes_73

import "testing"

// setZeroes 矩阵置为0
func setZeroes(matrix [][]int) {
	// 定义两个标记数组 防止首行首列被覆盖
	var m = make([]bool, len(matrix))
	var n = make([]bool, len(matrix[0]))
	// 遍历每一个元素
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// 对应首行,首列置为0
				matrix[i][0] = 0
				matrix[0][j] = 0
				m[i] = true
				n[j] = true
			}
		}
	}

	// 遍历首列 修改矩阵
	for i := 0; i < len(matrix); i++ {
		if m[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 遍历首行
	for i := 0; i < len(matrix[0]); i++ {
		if n[i] {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}

}

func TestName(t *testing.T) {
	// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	//var m = [][]int{{1,2,3}}
}

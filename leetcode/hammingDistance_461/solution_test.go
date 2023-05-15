package hammingDistance_461

import (
	"fmt"
	"testing"
)

// hammingDistance 异或算法 相同为0 不同为1
func hammingDistance(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor != 0 {
		if xor&1 == 1 {
			distance++
		}
		xor = xor >> 1
	}

	return distance
}

func TestName(t *testing.T) {
	fmt.Println(hammingDistance(1, 4))
}

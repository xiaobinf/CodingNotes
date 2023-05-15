package go_util

import "math/rand"

func RandInt() int {
	return rand.Int()
}

func RandIntN(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

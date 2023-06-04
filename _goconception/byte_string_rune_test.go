package _goconception

import (
	"fmt"
	"testing"
)

func TestByteStringRune(t *testing.T) {
	c := "我是谁中wss"
	fmt.Println(c, fmt.Sprintf("字符串长度为：%d", len(c)))
	fmt.Println([]byte(c), fmt.Sprintf("字节数组长度为：%d", len([]byte(c))))
	fmt.Println([]rune(c)[1:3], fmt.Sprintf("字符数组长度为：%d", len([]rune(c))))
	fmt.Println(string([]rune(c)[0:5]), fmt.Sprintf("字符数组长度为：%d", len([]rune(c))))
	// 定义数组
	var a [5]int
	fmt.Println(a)
	a[1] = 3
	fmt.Println(a)

	arr := [6]int{1, 2, 3, 4, 5} // 定义一个长度为 5 的整型数组，初始化为 {1, 2, 3, 4, 5}
	fmt.Println(arr)
	// byte 的底层类型是 uint8，而 rune 的底层类型是 int32
	var arr1 [5]byte // 定义长度为 5 的字节数组 byte 是 8 位的无符号整数类型（0 ~ 255)
	arr1[2] = 'a'
	var arr2 [10]rune // 定义长度为 10 的 rune 数组   rune 是 32 位的整数类型（0 ~ 4294967295）
	arr2[2] = '中'
	arr2[3] = 'a'
	var arr3 [2]string // 定义长度为 20 的字符串数组
	fmt.Println(arr1, arr2, arr3)

	// 定义一维列表

	// 定义二维列表

}

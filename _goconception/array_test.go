package _goconception

import (
	"fmt"
	"strings"
	"testing"
)

func TestArray(t *testing.T) {
	c := "我是谁中wss"
	fmt.Println(c, fmt.Sprintf("字符串长度为：%d", len(c)))

	var str = "123中文"
	fmt.Println(str[1], str[3:6])
	fmt.Println([]rune(str[3:6])) // 一个中文占3个字节

	fmt.Println(len([]rune("你好a")))

	var a = []int{1, 2, 3}
	fmt.Println(a)

	splitResult := strings.Split("12345", "1")
	fmt.Println(splitResult)
	fmt.Println(len(splitResult))

	// map操作
	m := make(map[int]int)
	m[1] = 2
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)
	delete(m, 4)
	fmt.Println(m)

}

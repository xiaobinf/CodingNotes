package _goconception

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	c := "我是谁中wss"
	fmt.Println(c, fmt.Sprintf("字符串长度为：%d", len(c)))

	var str = "123中文"
	fmt.Println(str[1], str[3:6])
	fmt.Println([]rune(str[3:6])) // 一个中文占3个字节

	fmt.Println(len([]rune("你好a")))
}

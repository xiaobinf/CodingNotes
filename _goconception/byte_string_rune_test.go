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
}

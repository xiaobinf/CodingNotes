package _goconception

import (
	"fmt"
	"testing"
)

// 测试单个文件
//go test -v interface_test.go

// 测试 单个方法
//go test -v interface_test.go -test.run TestName

// interface{} 统一万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("参数arg数据：")
	fmt.Println(arg)
}

// interface{} 如何判断是什么类型  "类型断言"机制  arg.(string)

func TestName(t *testing.T) {
	myFunc("hello world")
}

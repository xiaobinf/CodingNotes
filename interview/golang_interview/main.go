package main

import "fmt"

type A struct {
	s string
}

func f(s string) *A {
	a := new(A)
	a.s = s
	return a
}

func main() {
	a := f("hello")
	b := a.s + " world"
	c := b + "!"
	fmt.Println(c)
}

// 分析逃逸的编译命令：go build -gcflags "-m -l" main.go
//./main.go:9:8: leaking param: s
//./main.go:10:10: new(A) escapes to heap
//./main.go:17:11: a.s + " world" does not escape
//./main.go:18:9: b + "!" escapes to heap
//./main.go:19:13: ... argument does not escape
//./main.go:19:13: c escapes to heap

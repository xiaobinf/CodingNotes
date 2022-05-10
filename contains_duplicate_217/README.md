struct{}{}
空结构体不占用内存空间
package main

import (
"fmt"
"unsafe"
)

func main() {
fmt.Println(unsafe.Sizeof(struct{}{}))
}

Go 语言标准库没有提供 Set 的实现，通常使用 map 来代替。事实上，对于集合来说，只需要 map 的键，而不需要值。即使是将值设置为 bool 类型，也会多占据 1 个字节，那假设 map 中有一百万条数据，就会浪费 1MB 的空间。

因此呢，将 map 作为集合(Set)使用时，可以将值类型定义为空结构体，仅作为占位符使用即可。


type Set map[string]struct{}

func (s Set) Has(key string) bool {
_, ok := s[key]
return ok
}

func (s Set) Add(key string) {
s[key] = struct{}{}
}

func (s Set) Delete(key string) {
delete(s, key)
}

func main() {
s := make(Set)
s.Add("Tom")
s.Add("Sam")
fmt.Println(s.Has("Tom"))
fmt.Println(s.Has("Jack"))
}

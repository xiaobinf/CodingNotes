下面语法是否可以运行？
```go
type Param map[string]interface{}

type Show struct {
	Param
}

func main1() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
```
在这段代码中，Show 结构体嵌套了一个 Param 类型的匿名字段，因此 Show 类型的变量可以直接访问 Param 类型的字段。在 main1 函数中，我们创建了一个 Show 类型的变量 s，然后尝试给其 Param 字段中的 "RMB" 键赋值。

然而，由于 s 是一个指向 Show 类型的指针，而 s.Param 字段并没有被初始化，因此其值为 nil。在尝试给 s.Param["RMB"] 赋值时，会触发空指针异常，导致程序崩溃。

为了解决这个问题，我们需要在创建 Show 类型的变量时，先初始化其 Param 字段。可以使用 make 函数来创建一个空的 Param 类型的 map，然后将其赋值给 s.Param 字段。

```go
type Param map[string]interface{}

type Show struct {
    Param
}

func main() {
    s := &Show{Param: make(Param)}
    s.Param["RMB"] = 10000
}
```
在修改后的代码中，我们使用 make 函数来创建一个空的 Param 类型的 map，然后将其赋值给 s.Param 字段。这样，在尝试给 s.Param["RMB"] 赋值时，就不会触发空指针异常了。

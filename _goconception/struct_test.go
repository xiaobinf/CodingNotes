package _goconception

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Sum(num1 int32, num2 int32) int32 {
	return num1 + num2
}

// Person 定义对象的属性
type Person struct {
	name string
	age  int
}

// 定义对象的方法 本质上把方法绑在了对象上
// 注意 值传递
func (p Person) showName() {
	fmt.Println(p.name)
}

// 注意值传递
func (p Person) setName(name string) {
	p.name = name
}

// 指针传递才能改掉当前调用对象的值
func (p *Person) setPersonName(name string) {
	p.name = name
}

// SuperMan 继承对象
type SuperMan struct {
	Person
	level int
}

func (m SuperMan) showAbility() {
	fmt.Println("SuperMan info:")
	fmt.Println("name:", m.name)
	fmt.Println("age:", m.age)
	fmt.Println("level:", m.level)
}

func TestObject(t *testing.T) {
	p := Person{"fxb", 18}
	sp1 := SuperMan{p, 2}
	sp2 := SuperMan{Person{"zyt", 18}, 2}
	sp1.showAbility()
	sp2.showAbility()
	p.setPersonName("new name")
	p.setName("zzz")
	fmt.Println(p)
	assert.Equal(t, p.name, "new name")
	assert.Equal(t, sp1.name, "fxb")
	assert.Equal(t, sp1.level, 2)
	assert.Equal(t, sp2.name, "zyt")
	assert.Equal(t, sp2.level, 2)
}

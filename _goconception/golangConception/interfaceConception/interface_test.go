package interfaceConception

import (
	"fmt"
	"testing"
)

// 定义接口
type Bird interface {
	Fly()
	Type() string
}

type Crow struct {
	Name string
}

func (c *Crow) Fly() {
	fmt.Printf("我是%s，我用黄色的翅膀飞\n", c.Name)
}

func (c *Crow) Type() string {
	return c.Name
}

type Canary struct {
	Name string
}

func (c *Canary) Fly() {
	fmt.Printf("我是%s，我用黑色的翅膀飞\n", c.Name)
}

func (c *Canary) Type() string {
	return c.Name
}

func LetFly(bird Bird) {
	bird.Fly()
	fmt.Println(bird.Type())
}

func TestName(t *testing.T) {
	LetFly(&Crow{Name: "乌鸦"})
	LetFly(&Canary{Name: "金丝雀"})
}

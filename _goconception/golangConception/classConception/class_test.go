package classConception

import (
	"fmt"
	"testing"
)

type Bird struct {
	Type string
}

func (bird *Bird) Class() string {
	return bird.Type
}

// 定义接口
type Birds interface {
	Class() string
	Name() string
}

type Crow struct {
	Bird
	name string
}

func (c *Crow) Name() string {
	return c.name
}

func NewCrow(name string) *Crow {
	return &Crow{
		Bird: Bird{
			Type: "Crow",
		},
		name: name,
	}
}

type Canary struct {
	Bird
	name string
}

func (c *Canary) Name() string {
	return c.name
}

func NewCanary(name string) *Canary {
	return &Canary{
		Bird: Bird{
			Type: "Canary",
		},
		name: name,
	}
}

func BirdInfo(birds Birds) {
	fmt.Println(birds.Name(), birds.Class())
}

func TestName(t *testing.T) {
	canary := NewCanary("canaryA")
	crow := NewCanary("crowA")
	BirdInfo(canary)
	BirdInfo(crow)
}

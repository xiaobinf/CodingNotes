package goconception

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	var arr [2]int = [2]int{1, 2}
	var p *[2]int = &arr
	fmt.Println((*p)[0])
	fmt.Println(&p)
	assert.Equal(t, (*p)[0], 1)

	var a [2]int = [2]int{1, 2}
	var b [2]int = [2]int{1, 2}
	var p1 *[2]int = &a
	var p2 *[2]int = &b
	var p3 [2]*[2]int = [2]*[2]int{p1, p2}
	fmt.Println(p3)

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	p4, p5 := &arr1, &arr1[0]
	fmt.Println(p4)        //&[1,2,3,4,5] 整个数组的内存地址
	fmt.Println(p5)        //0xc0000180c0 数组第一个元素的内存地址
	fmt.Printf("%T\n", p4) //*[5]int    数组指针
	fmt.Printf("%T\n", p5) //*int        指针
}

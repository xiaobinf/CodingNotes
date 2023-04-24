package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)
}

func main() {
	s := Stack{
		data: []int{1, 2, 3},
	}
	s.push(1)
	s.push(2)
	fmt.Println(s.data)
}

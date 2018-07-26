package stack

import (
	"fmt"
)

type Stack struct {
	arr  []interface{}
	head int
}

func (s *Stack) Push(a interface{}) {
	s.arr = append(s.arr, a)
	s.head = s.head + 1
}

func (s *Stack) Pop() (interface{}, error) {
	if s.head >= 0 {
		a := s.arr[s.head]
		s.head = s.head - 1
		return a, nil
	}
	return nil, fmt.Errorf("Stack is empty")
}

func (s *Stack) Peek() (interface{}, error) {
	if s.head >= 0 {
		a := s.arr[s.head]
		return a, nil
	}
	return nil, fmt.Errorf("Stack is empty")
}

func NewStack() *Stack {
	s := &Stack{head: -1}
	return s
}

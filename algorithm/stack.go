package algorithm

import (
	"log"
)

type Stack struct {
	Size    int
	Stack   []int
	Pointer int
	Num     int
}

func NewStack(size int) *Stack {
	return &Stack{Size: size, Pointer: -1, Stack: make([]int, size)}
}

func (s *Stack) Push(item int) {
	if s.Num >= s.Size {
		log.Panic("Stack overflow")
	}
	s.Pointer += 1
	s.Num += 1
	s.Stack[s.Pointer] = item
}

func (s *Stack) Pop() int {
	if s.Num <= 0 {
		log.Panic("Stack empty")
	}
	item := s.Stack[s.Pointer]
	s.Pointer -= 1
	s.Num -= 1
	return item
}

func (s *Stack) IsEmpty() bool {
	if s.Num < s.Size {
		return true
	}
	return false
}

func (s *Stack) IsFull() bool {
	return !s.IsEmpty()
}

func (s *Stack) Top() int {
	if s.Num == 0 {
		log.Panic("Stack empty")
	}
	return s.Stack[s.Pointer]
}

func (s *Stack) GetSize() int {
	return s.Size
}

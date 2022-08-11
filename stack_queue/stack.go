package main

import (
	"errors"
)

type stack struct {
	value []int
}

func NewStack() *stack {
	s := stack{}
	return &s
}

func (s *stack) Push(n int) {
	(*s).value = append((*s).value, n)
}

func (s *stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}
	var r = (*s).value[len((*s).value)-1]
	(*s).value = (*s).value[:len((*s).value)-1]
	return r, nil
}

func (s *stack) Top() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}
	return (*s).value[len((*s).value)-1], nil
}

func (s *stack) Empty() bool {
	return len((*s).value) == 0
}

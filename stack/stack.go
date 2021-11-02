package stack

import (
	"fmt"
)

// Stack implementation
type Stack struct {
	data []float32
	size int
}

// NewStack instantiates a new stack
func NewStack(cap int) *Stack {
	return &Stack{data: make([]float32, 0, cap), size: 0}
}

// Push adds a new element at the end of the stack
func (s *Stack) Push(n float32) {
	s.data = append(s.data, n)
	s.size++
}

// Pop removes the last element from stack
func (s *Stack) Pop() (float32, error) {
	if s.IsEmpty() {
		return 0.0, fmt.Errorf("Stack is empty")
	}
	poped := s.Top()
	s.size--
	s.data = s.data[:s.size]
	return poped, nil
}

// Top returns the last element of stack
func (s *Stack) Top() float32 {
	return s.data[s.size-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

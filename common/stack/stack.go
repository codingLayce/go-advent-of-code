package stack

import (
	"fmt"
)

type Stack[T any] []T

// Push pushes on top of the stack
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// PushStack pushes the given stack on top on the stack
func (s *Stack[T]) PushStack(other Stack[T]) {
	*s = append(*s, other...)
}

// PushBack pushes in the bottom of the stack
func (s *Stack[T]) PushBack(value T) {
	tmp := *s
	*s = Stack[T]{value}
	*s = append(*s, tmp...)
}

// Pop removes and returns the top element on the stack
func (s *Stack[T]) Pop() T {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

// PopRange removes and returns the length top elements (in order)
func (s *Stack[T]) PopRange(length int) Stack[T] {
	toReturn := Stack[T]{}

	for i := 0; i < length; i++ {
		e := s.Pop()
		toReturn.PushBack(e)
	}

	return toReturn
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", *s)
}

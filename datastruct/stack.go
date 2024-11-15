package datastruct

import "fmt"

type Stack[T any] struct {
	slice []T
}

// Push adds an element to the stack by appending it to the slice.
func (s *Stack[T]) Push(elem T) {
	s.slice = append(s.slice, elem)
}

// Pop removes and return the last element of the slice.
func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	last := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return &last
}

// Peek looks at the top element without removing it.
func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}
	return &s.slice[len(s.slice)-1]
}

// IsEmpty checks if the slice is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// MoveOneElementTo moves one lement from current stack to another stack
func (s *Stack[T]) MoveOneElementTo(to *Stack[T]) error {
	topValue := s.Pop()
	if topValue == nil {
		return fmt.Errorf("can't move because stack is empty")
	}
	to.Push(*topValue)
	return nil
}

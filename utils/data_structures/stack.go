package data_structures

import "errors"

type StackInterface[T any] interface {
	Push(value T) error
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	GetSize() int
}

type Stack[T any] struct {
	Elements []T
	Top      int
	Size     int
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Top == 0
}

func (s *Stack[T]) GetSize() int {
	return s.Top
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.Elements[s.Top-1], nil
}

func (s *Stack[T]) Push(value T) error {
	if s.Top >= s.Size {
		return errors.New("stack overflow")
	}
	s.Elements = append(s.Elements, value)
	s.Top++
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	s.Top--
	value := s.Elements[s.Top]
	s.Elements = s.Elements[:s.Top]
	return value, nil
}

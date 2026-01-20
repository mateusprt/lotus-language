package utils

import "github.com/mateusprt/lox-language/utils/data_structures"

func NewStack[T any](size int) *data_structures.Stack[T] {
	return &data_structures.Stack[T]{
		Elements: make([]T, 0, size),
		Top:      0,
		Size:     size,
	}
}

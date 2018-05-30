package main

import "errors"

// Stack data structure
type Stack struct {
	elements []interface{}
	size     int
}

// CreateStack function creates a new stack instance
func CreateStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

// Push an element into the stack s
func (s *Stack) Push(e interface{}) {
	// Add the new element e at the end of the slice
	s.elements = append(s.elements, e)
	s.size++
}

// Pop the top element from the stack s
func (s *Stack) Pop() (interface{}, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("No more elements in this stack")
	}

	top := s.elements[len(s.elements)-1]

	// remove the top element
	s.elements = s.elements[0 : len(s.elements)-1]

	// Decrease stack size
	s.size--
	return top, nil
}

// IsEmpty method check if the stack s contains any element
func (s *Stack) isEmpty() bool {
	if s == nil {
		return true
	}
	return len(s.elements) == 0
}

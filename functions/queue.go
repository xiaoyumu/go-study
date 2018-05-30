package main

import (
	"errors"
)

// Queue data structure
type Queue struct {
	elements []interface{}
	size     int
}

// CreateQueue function creates a new queue instance
func CreateQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func (q *Queue) enqueue(e interface{}) {
	q.elements = append(q.elements, e)
	q.size++

}

func (q *Queue) dequeue() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("No more elements in this queue")
	}

	head := q.elements[0]
	q.elements = q.elements[1:]
	q.size--

	return head, nil
}

func (q *Queue) isEmpty() bool {
	if q == nil {
		return true
	}
	return len(q.elements) == 0
}

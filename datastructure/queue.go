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

// Enqueue puts an element into the queue tail
func (q *Queue) Enqueue(e interface{}) {
	q.elements = append(q.elements, e)
	q.size++

}

// Dequeue gets an element from the queue head, and remove it from queue
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("No more elements in this queue")
	}

	head := q.elements[0]
	q.elements = q.elements[1:]
	q.size--

	return head, nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return len(q.elements) == 0
}

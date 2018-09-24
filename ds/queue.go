package ds

import (
	"errors"
	// "github.com/victor-iyiola/go-algorithms/ds"
)

// Data is the data type for elements inside Queue.
// type QueueData float64

// Queue data structure obeys the First In First Out (FIFO) lingo.
type Queue struct {
	data  []Data
	count uint
}

// NewQueue creates a new Queue instance.
func NewQueue() *Queue {
	return &Queue{data: nil, count: 0}
}

// Push adds new element to the back of a queue.
func (q *Queue) Push(data Data) {
	q.data = append(q.data, data)
	q.count = uint(len(q.data))
}

// Peek returns the first element in the Queue.
func (q *Queue) Peek() (Data, error) {
	if q.count > 0 {
		return q.data[0], nil
	}

	return 0, errors.New("Queue Empty: Queue is empty")
}

// Pop removes the first element from the Queue.
func (q *Queue) Pop() (Data, error) {
	// data to be returned.
	var data Data

	// No elements in the queue.
	if q.Empty() {
		return data, errors.New("QueueError: No element in the queue")
	}

	// Get the first element in the slice.
	data, _ = q.Peek()

	// Delete the last element from the queue.
	q.data = q.data[0 : q.count-1]

	// Decrement count.
	q.count = uint(len(q.data))

	return data, nil
}

// Count returns the total number of elements in the queue.
func (q *Queue) Count() int {
	return int(q.count)
}

// Empty indicates whether or not the Queue has any elements in it.
func (q *Queue) Empty() bool {
	return q.count <= 0
}

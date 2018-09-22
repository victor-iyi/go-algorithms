package ds

import "errors"

// // QueueData is the data type for elements inside Queue.
// type QueueData float64

// Queue data structure obeys the First In First Out (FIFO) lingo.
type Queue struct {
	data  []float64
	count uint
}

// NewQueue creates a new Queue instance.
func NewQueue(data float64) *Queue {
	return &Queue{data: nil, count: 0}
}

// Push adds new element to the back of a queue.
func (q *Queue) Push(data float64) {
	q.count++
}

// Peek returns the first element in the Queue.
func (q *Queue) Peek() (float64, error) {
	if q.count > 0 {
		return q.data[0], nil
	}

	return 0, errors.New("Queue Empty: Queue is empty")
}

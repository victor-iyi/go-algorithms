package ds

// QueueData is the data type for elements inside Queue.
type QueueData string

// Queue data structure obeys the First In First Out (FIFO) lingo.
type Queue struct {
	data  []Queue
	count uint
}

// NewQueue creates a new Queue instance.
func NewQueue(data QueueData) *Queue {
	return &Queue{data: nil, count: 0}
}

// Push adds new element to the back of a queue.
func (q *Queue) Push(data QueueData) {
	q.count++
}

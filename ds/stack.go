package ds

import "errors"

// Data contains data in a stack.
// type StackData float64

// Stack is a data structure that abides to the Last In First Out (LIFO) lingo.
type Stack struct {
	data  []Data // Contains stack data.
	count uint   // An internal element count maintained by stack.
}

// NewStack creates a new instance of a stack.
func NewStack() *Stack {
	return &Stack{data: nil, count: 0}
}

// Peek returns the last element in the stack.
func (s *Stack) Peek() (Data, error) {
	if s.Empty() {
		return 0, errors.New("StackError: Stack is empty")
	}

	return s.data[s.count-1], nil
}

// Pop removes the last element from the stack.
func (s *Stack) Pop() (Data, error) {
	// data to be returned.
	var data Data

	// No elements in the stack.
	if s.Empty() {
		return data, errors.New("StackError: No element in the stack")
	}

	// Get the last element in the slice.
	data, _ = s.Peek()

	// Delete the last element from the stack.
	s.data = s.data[:s.count]

	// Decrement count.
	s.count--

	return data, nil
}

// Push adds a new element to the end of a stack.
func (s *Stack) Push(data Data) {
	// Adds new data into the stack.
	s.data = append(s.data, data)

	// Increment count after successful push.
	s.count = uint(len(s.data))
}

// Count returns the number of elements in a stack.
func (s *Stack) Count() int {
	return int(s.count)
}

// Empty indicates whether the stack is empty or not.
// Returns `true` if it's empty and `false` otherwise.
func (s *Stack) Empty() bool {
	return s.count == 0
}

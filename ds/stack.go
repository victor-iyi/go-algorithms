package ds

import "errors"

// StackData contains data in a stack.
type StackData float64

// Stack is a data structure that abides to the Last In First Out (LIFO) lingo.
type Stack struct {
	data  []StackData // Contains stack data.
	count uint        // An internal element count maintained by stack.
}

// NewStack creates a new instance of a stack.
func NewStack() *Stack {
	return &Stack{data: nil, count: 0}
}

// Peek returns the last element in the stack.
func (s *Stack) Peek() (StackData, error) {
	if s.count == 0 {
		return 0, errors.New("StackError: Stack is empty")
	}

	return s.data[s.count], nil
}

// Pop removes the last element from the stack.
func (s *Stack) Pop() {
	if s.count == 0 {
		// delete(s.data, s.count)
	}
	// Decrement count.
	s.count--
}

// Push adds a new element to the end of a stack.
func (s *Stack) Push(data StackData) {
	// Adds new data into the stack.
	s.data = append(s.data, data)

	// Increment count after successful push.
	s.count = uint(len(s.data))
}

// Count returns the number of elements in a stack.
func (s *Stack) Count() uint {
	return s.count
}

// Empty indicates whether the stack is empty or not.
// Returns `true` if it's empty and `false` otherwise.
func (s *Stack) Empty() bool {
	return s.count == 0
}

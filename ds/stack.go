// Package ds implements Stack data structure in Go.
package ds

// Stack is a data structure that abides to the Last In Last Out (LIFO) lingo.
type Stack struct {
	// An internal element count maintained by stack.
	count uint
}

// Peek returns the last element in the stack.
func (s *Stack) Peek() {

}

// Pop removes the last element from the stack.
func (s *Stack) Pop() {

}

// Push adds a new element to the end of a stack.
func (s *Stack) Push() {

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

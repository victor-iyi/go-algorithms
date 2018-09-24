package main

import (
	"fmt"

	"github.com/victor-iyiola/go-algorithms/ds"
)

func main() {
	stack := ds.NewStack()
	stack.Push(3.14)
	stack.Push(13.)
	stack.Push(.21)
	stack.Push(526.1)
	stack.Push(190.21)
	printDetails(stack)
}

func printDetails(s *ds.Stack) {
	// Print empty status.
	fmt.Print("Stack is ")
	if s.Empty() {
		fmt.Println("empty")
		return // End the function if stack is empty.
	}
	fmt.Println("not empty")

	// Number of elements contained in a stack.
	fmt.Printf("Stack contains: %d element(s)\n", s.Count())

	// Peek into the stack.
	peek, err := s.Peek()
	if err != nil {
		fmt.Printf("Cannot peek [%s]\n", err)
	} else {
		fmt.Println("s.Peek() =", peek)
	}

	// Pop last element from stack.
	var pop ds.StackData
	for i := 0; i < int(s.Count()); i++ {
		pop, err = s.Pop()

		if err != nil {
			fmt.Printf("Cannot pop [%s]\n", err)
		} else {
			fmt.Println("s.Pop() =", pop)
		}
	}

}

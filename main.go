package main

import (
	"fmt"
	"strings"

	"github.com/victor-iyiola/go-algorithms/ds"
)

func main() {
	// Trie example.
	trie := ds.NewTrieFromChar('*')
	usingTrie(trie)

	divider()

	// Stack example.
	stack := ds.NewStack()
	usingStack(stack)

	divider()

	// Queue example.
	queue := ds.NewQueue()
	usingQueue(queue)
}

func divider(count ...int) {
	if count == nil {
		count = append(count, 50)
	}

	fmt.Printf("\n%s\n\n", strings.Repeat("=", count[0]))
}

func usingStack(s *ds.Stack) {
	// Push elements to the stack.
	s.Push(3.14)
	s.Push(13.)
	s.Push(.21)
	s.Push(526.1)
	s.Push(190.21)

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
	for i := 0; i < s.Count(); i++ {
		pop, err = s.Pop()

		if err != nil {
			fmt.Printf("Cannot pop [%s]\n", err)
		} else {
			fmt.Println("s.Pop() =", pop)
		}
	}

}

func usingQueue(q *ds.Queue) {
	// Push elements to the queue.
	q.Push(3.14)
	q.Push(13.)
	q.Push(.21)
	q.Push(526.1)
	q.Push(190.21)

	// Print empty status.
	fmt.Print("Queue is ")
	if q.Empty() {
		fmt.Println("empty")
		return // End the function if queue is empty.
	}
	fmt.Println("not empty")

	// Number of elements contained in a queue.
	fmt.Printf("Queue contains: %d element(s)\n", q.Count())

	// Peek into the queue.
	peek, err := q.Peek()
	if err != nil {
		fmt.Printf("Cannot peek [%s]\n", err)
	} else {
		fmt.Println("q.Peek() =", peek)
	}

	// Pop last element from queue.
	var pop ds.QueueData
	for i := 0; i < q.Count(); i++ {
		pop, err = q.Pop()

		if err != nil {
			fmt.Printf("Cannot pop [%s]\n", err)
		} else {
			fmt.Println("q.Pop() =", pop)
		}
	}

}

func usingTrie(t *ds.Trie) {
	// Add single word to the trie.
	t.Add("bell")
	t.Add("base")
	t.Add("beat")
	t.Add("bellman")

	// Add multiple words to the trie.
	t.AddAll("victor victoria victory victories")
	t.AddAll("programming terminal commit async escape milk terminate program")

	// Print trie contents.
	ds.Print(t.Root)

	// Check occurrence of some word prefix in the trie.
	hasVic, countVic := t.Find("vic")
	hasProg := t.Contains("prog")

	// Print results.
	fmt.Printf("root.Char = %c\n", t.Root.Char)
	fmt.Printf("hasVic = %t\tcountVic = %d\n", hasVic, countVic)
	fmt.Printf("hasProg = %t\n", hasProg)
}

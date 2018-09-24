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
	usingLinearDS("Stack", stack)

	divider()

	// Queue example.
	queue := ds.NewQueue()
	usingLinearDS("Queue", queue)
}

func divider(count ...int) {
	if count == nil {
		count = append(count, 50)
	}

	fmt.Printf("\n%s\n\n", strings.Repeat("=", count[0]))
}

func usingLinearDS(name string, linear ds.LinearDS) {
	// Push elements to the dataStructure.
	linear.Push(3.14)
	linear.Push(13.)
	linear.Push(.21)
	linear.Push(526.1)
	linear.Push(190.21)

	// Print empty status.
	fmt.Printf("%s is ", name)
	if linear.Empty() {
		fmt.Println("empty")
		return // End the function if dataStructure is empty.
	}
	fmt.Println("not empty")

	// Number of elements contained in a dataStructure.
	fmt.Printf("%s contains: %d element(s)\n", name, linear.Count())

	// Peek into the dataStructure.
	peek, err := linear.Peek()
	if err != nil {
		fmt.Printf("Cannot peek [%s]\n", err)
	} else {
		fmt.Println("dataStructure.Peek() =", peek)
	}

	// Pop last element from stack.
	for i := 0; i < linear.Count(); i++ {
		pop, err := linear.Pop()

		if err != nil {
			fmt.Printf("Cannot pop [%s]\n", err)
		} else {
			fmt.Println("dataStructure.Pop() =", pop)
		}
	}
}

func usingTrie(t *ds.Trie) {
	// Add single word to the trie.
	t.Push("bell")
	t.Push("base")
	t.Push("beat")
	t.Push("bellman")

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

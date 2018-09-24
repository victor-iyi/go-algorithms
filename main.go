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
	usingDataStructure("Stack", stack)

	divider()

	// Queue example.
	queue := ds.NewQueue()
	usingDataStructure("Queue", queue)
}

func divider(count ...int) {
	if count == nil {
		count = append(count, 50)
	}

	fmt.Printf("\n%s\n\n", strings.Repeat("=", count[0]))
}

func usingDataStructure(name string, dataStructure ds.DataStructure) {
	// Push elements to the dataStructure.
	dataStructure.Push(3.14)
	dataStructure.Push(13.)
	dataStructure.Push(.21)
	dataStructure.Push(526.1)
	dataStructure.Push(190.21)

	// Print empty status.
	fmt.Printf("%s is ", name)
	if dataStructure.Empty() {
		fmt.Println("empty")
		return // End the function if dataStructure is empty.
	}
	fmt.Println("not empty")

	// Number of elements contained in a dataStructure.
	fmt.Printf("%s contains: %d element(s)\n", name, dataStructure.Count())

	// Peek into the dataStructure.
	peek, err := dataStructure.Peek()
	if err != nil {
		fmt.Printf("Cannot peek [%s]\n", err)
	} else {
		fmt.Println("dataStructure.Peek() =", peek)
	}

	// Pop last element from stack.
	for i := 0; i < dataStructure.Count(); i++ {
		pop, err := dataStructure.Pop()

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

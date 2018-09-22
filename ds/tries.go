package ds

import (
	"fmt"
	"strings"
)

// TrieNode represents each node in a Trie.
type TrieNode struct {
	Char     rune
	Children []*TrieNode
	IsWord   bool
	Count    uint8
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode(char rune) *TrieNode {
	return &TrieNode{Char: char, Children: nil, IsWord: false, Count: 0}
}

// AddChild adds a new TrieNode to it's children.
func (tn *TrieNode) AddChild(node *TrieNode) {
	tn.Children = append(tn.Children, node)
}

// Empty return false if node has no children.
func (tn *TrieNode) Empty() bool {
	return (len(tn.Children) <= 0)
}

// Trie is a tree data structures that consist of nodes of characters
// which forms a valid word as you traverse down the tree.
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new trie with initialized root.
func NewTrie(node *TrieNode) *Trie {
	return &Trie{root: node}
}

// Add adds a new Node to a the trie.
func (t *Trie) Add(word string) {
	// Start at the root node.
	node := t.root

	// Loop through each character in the given word.
	for _, char := range word {
		// Has char been found in the Trie?
		var found bool

		// Go through current node's children.
		for _, child := range node.Children {
			// If the current `node.char` == `char`,
			// increment node count
			if child.Char == char {
				// Found character in the Trie.
				found = true

				// Increment possible words counter & update next node.
				node.Count++
				node = child

				break
			}
		}

		// If the char was not found.
		if !found {
			// Add child node to current node.
			newNode := NewTrieNode(char)
			node.AddChild(newNode)

			// Update the node.
			node = newNode
		}
	}

	// After adding all characters to the Trie, mark the last char as a complete word.
	node.IsWord = true
}

// AddAll adds multiple words in a sentence to the trie.
func (t *Trie) AddAll(sentence string) {
	for _, word := range strings.Fields(sentence) {
		t.Add(word)
	}
}

// Find returns whether or not a prefix has been found,
// and also how many words it can make if found.
func (t *Trie) Find(prefix string) (bool, uint8) {
	// Start from the root node.
	node := t.root

	// Return false if there's no children in root node.
	if len(node.Children) < 1 {
		return false, 0
	}

	for _, char := range prefix {
		var found bool
		for _, child := range node.Children {
			if child.Char == char {
				found = true
				node = child
				break
			}
		}

		if !found {
			return false, 0
		}
	}

	return true, node.Count
}

// Contains returns if a prefix contains a given prefix.
func (t *Trie) Contains(prefix string) bool {
	found, _ := t.Find(prefix)
	return found
}

// Print prints out the children in a given trie.
func Print(tn *TrieNode) {
	for _, child := range tn.Children {
		if len(child.Children) == 0 {
			return
		}
		fmt.Printf("%c ", child.Char)
		Print(child)
	}
}

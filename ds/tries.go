package ds

import "fmt"

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
	node := t.root

	for _, char := range word {
		var found bool

		for _, child := range node.Children {
			if child.Char == char {
				found = true
				node.Count++
				node = child
				break
			}
		}

		// If the char was not found.
		if !found {
			newNode := NewTrieNode(char)
			node.AddChild(newNode)
			node = newNode
		}
	}

	node.IsWord = true
}

// AddAll adds multiple words in a sentence to the trie.
func (t *Trie) AddAll(sentence string) {

}

// Find returns whether or not a prefix has been found,
// and also how many words it can make if found.
func (t *Trie) Find(prefix string) (bool, uint8) {
	node := t.root

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

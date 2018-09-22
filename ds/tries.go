package ds

// TrieNode represents each node in a Trie.
type TrieNode struct {
	Data     string
	Children []TrieNode
	IsWord   bool
	Count    uint8
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode(data string) *TrieNode {
	return &TrieNode{Data: data, Children: nil, IsWord: false, Count: 0}
}

// AddChild adds a new TrieNode to it's children.
func (node *TrieNode) AddChild() {

}

// Trie is a tree data structures that consist of nodes of characters
// which forms a valid word as you traverse down the tree.
type Trie struct {
	root TrieNode
}

// NewTrie creates a new trie with initialized root.
func NewTrie(node *TrieNode) *Trie {
	return &Trie{root: *node}
}

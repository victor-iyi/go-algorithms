package main

import (
	"fmt"

	"github.com/victor-iyiola/go-algorithms/ds"
)

func main() {
	root := ds.NewTrieNode('*')
	trie := ds.NewTrie(root)

	trie.Add("victor")
	trie.Add("victoria")
	trie.Add("bell")
	trie.Add("base")
	trie.Add("beat")
	trie.Add("bellman")

	ds.Print(root)

	hasVic, countVic := trie.Find("vic")
	hasProg := trie.Contains("prog")

	fmt.Printf("root.Char = %c\n", root.Char)
	fmt.Printf("hasVic = %t\tcountVic = %d\n", hasVic, countVic)
	fmt.Printf("hasProg = %t\n", hasProg)
}

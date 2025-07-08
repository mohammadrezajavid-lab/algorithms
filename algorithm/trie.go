package algorithm

import "log"

type TrieNode struct {
	mark  bool
	char  rune
	edges []*TrieNode
}

type Trie struct {
	root *TrieNode
}

func newNode(char rune) *TrieNode {
	return &TrieNode{mark: false, char: char, edges: make([]*TrieNode, 26)}
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{
		mark:  false,
		char:  0,
		edges: make([]*TrieNode, 26),
	}}
}

func (trie *Trie) GetRoot() *TrieNode {
	return trie.root
}

// در این صورت اوردر درج میشود به اندازه ی طول رشته ای که میخواهد درج شود مثلا اکه طول رشته k باشد میشود O(k)
func (trie *Trie) Insert(str []rune, node *TrieNode, idx int) {
	if idx == len(str) {
		node.mark = true
		return
	}

	index := str[idx] - 'a'
	if node.edges[index] == nil {
		node.edges[index] = newNode(str[idx])
	}

	trie.Insert(str, node.edges[index], idx+1)
}

// همچنین هم با O(k) رشته را پیدا میکند
func (trie *Trie) Find(str []rune, node *TrieNode, idx int) bool {
	if idx == len(str) {
		return node.mark
	}

	index := str[idx] - 'a'
	if node.edges[index] == nil {
		return false
	}

	return trie.Find(str, node.edges[index], idx+1)
}

// و همچنین با O(k) هم رشته با delete میکند
func (trie *Trie) Delete(str []rune, node *TrieNode, idx int) {
	if idx == len(str) {
		if node.mark == false {
			log.Fatal("Item not found")
		}
		node.mark = false
		return
	}

	index := str[idx] - 'a'
	if node.edges[index] == nil {
		log.Fatal("Item not found2")
	}

	trie.Delete(str, node.edges[index], idx+1)
}

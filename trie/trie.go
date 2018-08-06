package trie

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	nodes []*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	nodes := make([]*TrieNode, 26)
	node := &TrieNode{nodes: nodes, isEnd: false}
	return node
}

func (root *TrieNode) Insert(value string) {
	valueRune := []rune(strings.ToLower(value))
	aValue := []rune("a")[0]
	for _, v := range valueRune {
		index := v - aValue
		if root.nodes[index] == nil {
			root.nodes[index] = NewTrieNode()
		}
		root = root.nodes[index]
	}
	root.isEnd = true
}

func (root *TrieNode) Delete(value string) {
	aValue := []rune("a")[0]
	if root.Find(value) {
		valueRune := []rune(strings.ToLower(value))
		root.nodes[valueRune[0]-aValue].deleteUtil(valueRune, root)
	} else {
		fmt.Println("Value not found")
	}
}

func (root *TrieNode) deleteUtil(valueRune []rune, parent *TrieNode) {
	aValue := []rune("a")[0]
	if len(valueRune) > 1 {
		root.nodes[valueRune[1]-aValue].deleteUtil(valueRune[1:], root)
	}
	flag := false
	for _, n := range root.nodes {
		if n != nil {
			flag = true
		}
	}
	if flag {
		root.isEnd = false
	} else {
		parent.nodes[valueRune[0]-aValue] = nil
	}
}

func (root *TrieNode) Find(value string) bool {
	valueRune := []rune(strings.ToLower(value))
	aValue := []rune("a")[0]
	for _, v := range valueRune {
		index := v - aValue
		if root.nodes[index] == nil {
			return false
		}
		root = root.nodes[index]
	}
	return root.isEnd
}

func (root *TrieNode) Print() {
	aValue := []rune("a")[0]
	for i := 0; i < len(root.nodes); i++ {
		if root.nodes[i] != nil {
			fmt.Println(string(rune(i)+aValue), root.nodes[i].isEnd)
			root.nodes[i].Print()
		}
	}
}

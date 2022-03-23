package main

type Trie struct {
	rootNode *Node
}

type Node struct {
	children map[string]*Node
	isEnd    bool
}

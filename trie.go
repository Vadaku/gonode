package main

//Each non leaf node has a pointer to another node.

type Trie struct {
	rootNode  *Node
	rotations map[string]*Node
}

type Node struct {
	children map[byte]*Node
	childArr []string
	isEnd    bool
}

//Init a Trie with an empty root node.
//Initialize the Nodes children field
func initializeTrie() *Trie {
	initTrie := &Trie{rootNode: &Node{}}
	initTrie.rootNode.children = make(map[byte]*Node)
	return initTrie
}

//Create new node and inititialize the Nodes children field
func addChildNode() *Node {
	newNode := &Node{}
	newNode.children = make(map[byte]*Node)
	return newNode
}

// Iterate through each node in the Trie checking if the nodes char matches params char.
// If not, create a new node and add the param char to this node, then add this node to the trie.
func (t *Trie) insertToTrie(w string, rotation string) {
	//Split the parameter word string into a slice of bytes
	keySlice := []byte(w)
	currentNode := t.rootNode
	for i := 0; i < len(keySlice); i++ {
		//If child node is empty then create a new empty node.
		if _, ok := currentNode.children[keySlice[i]]; !ok {
			currentNode.children[keySlice[i]] = addChildNode()
		}
		currentNode = currentNode.children[keySlice[i]]
	}
	currentNode.childArr = append(currentNode.childArr, rotation)
	currentNode.isEnd = true
}

//Search if the Trie includes the given prefix.
//Returns two values, result slice and bool if the Trie contains the prefix.
//Returns nil result if prefix is not in the trie.
func (t *Trie) searchTrie(key string) ([]string, bool) {
	currentNode := t.rootNode
	keySlice := []byte(key)
	for i := 0; i < len(keySlice); i++ {
		if _, ok := currentNode.children[keySlice[i]]; !ok {
			return nil, false
		}
		currentNode = currentNode.children[keySlice[i]]
	}
	return currentNode.childArr, currentNode.isEnd
}

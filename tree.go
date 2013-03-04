package tree

import "fmt"

type Tree struct {
	root *Node
	size int
}

type Node struct {
	left, right *Node
	key string
	value string
}

func Create() *Tree {
	t := new(Tree)
	t.size = 0
	return t
}

func (t *Tree) Insert(key string, value string) {
		t.root = t.root.insertHelper(key, value)

		t.size++
}

func (node *Node) insertHelper(key string, value string) *Node {
	if(node == nil) {
		node = new(Node)
		node.key = key
		node.value = value
	}

	if(node.key > key) {
		fmt.Println(key, "less than", node.key)
		node.left = node.left.insertHelper(key, value)
	}

	if(node.key < key) {
		fmt.Println(key, "greater than", node.key)
		node.right = node.right.insertHelper(key, value)
	}

	return node
}

func (t *Tree) Get(key string) string {
	if t.size == 0 {
		return ""
	}

	return t.root.getHelper(key)
}


func (node *Node) getHelper(key string) string {
	if(node == nil) {
		return ""
	}

	if(node.key == key) {
		return node.value
	}

	if(node.key > key) {
		return node.left.getHelper(key)
	}

	if(node.key < key) {
		return node.right.getHelper(key)
	}

	return ""
}

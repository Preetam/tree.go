package tree

import (
	"fmt"
	"errors"
)

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
		t.root, _ = t.root.insertHelper(key, value)
		t.size++
}

func (node *Node) insertHelper(key string, value string) (*Node, error) {
	if(node == nil) {
		node = new(Node)
		node.key = key
		node.value = value
	}

	if(node.key > key) {
		fmt.Println(key, "less than", node.key)
		node.left, _ = node.left.insertHelper(key, value)
	}

	if(node.key < key) {
		fmt.Println(key, "greater than", node.key)
		node.right, _ = node.right.insertHelper(key, value)
	}

	return node, nil
}

func (t *Tree) Get(key string) (string, error) {
	if t.size == 0 {
		return "", errors.New("Empty tree");
	}

	return t.root.getHelper(key), nil
}


func (node *Node) getHelper(key string) string {
	if(node.key == key) {
		return node.value
	}

	if(node.key > key) {
		return node.left.getHelper(key)
	}

	if(node.key < key) {
		return node.right.getHelper(key)
	}

	return node.value
}

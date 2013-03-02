package tree

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
	if t.root == nil {
		t.root = new(Node)
		t.root.key = key
		t.root.value = value
		t.size++
	}
}

func (node *Node) insertHelper(key string, value string) {
	if(node == nil) {
		node = new(Node)
		node.key = key
		node.value = value
	}

	if(node.key > key) {
		node.left.insertHelper(key, value)
	}

	if(node.key < key) {
		node.right.insertHelper(key, value)
	}
}

func (t *Tree) Get(key string) string {
	if t.size == 0 {
		return ""
	}

	if t.size == 1 {
		if(t.root.key == key) {
			return t.root.value
		}

		if(t.root.key != key) {
			return ""
		}
	}

	return ""
}


func (node *Node) getHelper(key string) string {
	if(node == nil) {
		return ""
	}

	if(node.key == key) {
		return node.value
	}

	if(node.key < key) {
		return node.left.getHelper(key)
	}

	if(node.key > key) {
		return node.right.getHelper(key)
	}

	return ""
}

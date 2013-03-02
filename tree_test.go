package tree

import (
	"testing"
)

func Test1(t *testing.T) {
	tree := Create()
	if tree.size != 0 {
		t.Errorf("blah")
	}
}

func Test2(t *testing.T) {
	tree := Create()
	tree.Insert("a", "val")
	if tree.size != 1 {
		t.Errorf("blah")
	}
}

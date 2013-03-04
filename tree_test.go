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
	val := tree.Get("a")

	if tree.size != 1 {
		t.Errorf("blah")
	}

	if val != "val" {
		t.Errorf("blah")
	}
}

func Test3(t *testing.T) {
	tree := Create()
	tree.Insert("a", "A")
	tree.Insert("b", "B")
	tree.Insert("c", "C")
	valA := tree.Get("a")
	valB := tree.Get("b")
	valC := tree.Get("c")

	if tree.size != 3 {
		t.Errorf("blah %v", tree.size)
	}

	if valA != "A" {
		t.Errorf("blah %v", valA)
	}
	if valB != "B" {
		t.Errorf("blah %v", valB)
	}
	if valC != "C" {
		t.Errorf("blah %v", valC)
	}
}

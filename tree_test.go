package tree

import (
	"testing"
	"fmt"
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
	val, err := tree.Get("a")

	if(err != nil) {
		t.Error(err)
	}

	if tree.size != 1 {
		t.Errorf("blah")
	}

	if val != "val" {
		t.Errorf("blah")
	}
}

func Test3(t *testing.T) {
	tree := Create()
	tree.Insert("b", "B")
	tree.Insert("c", "C")
	tree.Insert("a", "A")
	valA, errA := tree.Get("a")
	valB, errB := tree.Get("b")
	valC, errC := tree.Get("c")


	if(errA != nil) {
		t.Error(errA)
	}
	if(errB != nil) {
		t.Error(errB)
	}
	if(errC != nil) {
		t.Error(errC)
	}

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

	fmt.Println(tree.root.left.value)
}
package ds

import "testing"

func TestTree(t *testing.T) {
	tree := NewTree(1)
	if tree.Val != 1 {
		t.Error("new a tree with root value 1")
	}

	l := tree.Insert(2, true)
	if l.Val != 2 {
		t.Error("insert left leaf 2")
	}

	r := tree.Insert(3, false)

	if r.Val != 3 {
		t.Error("insert right leaf 3")
	}

	tree.Desc()
}

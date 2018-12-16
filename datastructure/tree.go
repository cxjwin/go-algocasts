package ds

import "fmt"

// Tree - tree struct
type Tree struct {
	Left  *Tree
	Value interface{}
	Right *Tree
}

// Insert - insert node to tree
func (t *Tree) Insert(v interface{}, left bool) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if left {
		t.Left = t.Left.Insert(v, true)
		return t.Left
	}

	t.Right = t.Right.Insert(v, false)
	return t.Right
}

// Desc - description of tree
func (t *Tree) Desc() {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	t.Left.Desc()
	t.Right.Desc()
}

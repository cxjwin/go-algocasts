package algo

import "fmt"

// Tree - tree struct
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) insert(v int, left bool) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if left {
		t.Left = t.Left.insert(v, true)
		return t.Left
	}

	t.Right = t.Right.insert(v, false)
	return t.Right
}

func (t *Tree) description() {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	t.Left.description()
	t.Right.description()
}

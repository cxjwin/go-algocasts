package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

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

func isSymmetric(l *Tree, r *Tree) bool {
	if l != nil && r != nil {
		return l.Value == r.Value &&
			isSymmetric(l.Left, r.Right) &&
			isSymmetric(l.Right, r.Left)
	}

	return l == nil && r == nil
}

func isSymmetricTreeRecursive(t *Tree) bool {
	if t == nil {
		return true
	}

	return isSymmetric(t.Left, t.Right)
}

func isSymmetricTreeIterative(t *Tree) bool {
	if t == nil {
		return true
	}

	s := new(stack.Stack)
	s.Push(t.Left)
	s.Push(t.Right)

	for s.Len() != 0 {
		l := s.Pop()
		r := s.Pop()
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		tl := l.(*Tree)
		tr := r.(*Tree)
		if tl.Value != tr.Value {
			return false
		}
		if tl.Left != nil {
			s.Push(tl.Left)
		}
		if tr.Right != nil {
			s.Push(tr.Right)
		}
		if tl.Right != nil {
			s.Push(tl.Right)
		}
		if tr.Left != nil {
			s.Push(tr.Left)
		}
	}

	return true
}

func main() {
	root := Tree{nil, 1, nil}

	// left
	l1 := root.insert(2, true)
	l1.insert(3, true)
	l1.insert(4, false)

	// right
	l2 := root.insert(2, false)
	l2.insert(4, true)
	l2.insert(3, false)

	// print
	root.description()

	// check
	fmt.Printf("recursive - tree is symmetric %v\n", isSymmetricTreeRecursive(&root))
	fmt.Printf("iterative - tree is symmetric %v\n", isSymmetricTreeIterative(&root))
}

package ds

import (
	"testing"
)

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

func TestBuildTree(t *testing.T) {
	nums := []interface{}{3, 2, 3, nil, 3, nil, 1}
	root := BuildTree(nums)

	//      3
	//     / \
	//    2   3
	//     \   \
	//      3   1
	if root.Val != 3 ||
		root.Left.Val != 2 ||
		root.Left.Right.Val != 3 ||
		root.Right.Val != 3 ||
		root.Right.Right.Val != 1 {
		t.Error("1 - build error")
	}

	nums = []interface{}{1, nil, 2, nil, 4}
	root = BuildTree(nums)
	//      1
	//       \
	//        2
	//         \
	//          4
	if root.Val != 1 ||
		root.Right.Val != 2 ||
		root.Right.Right.Val != 4 {
		t.Error("2 - build error")
	}
}

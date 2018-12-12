package algo

import (
	"fmt"
	"testing"
)

func TestSymmetricTree(t *testing.T) {
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

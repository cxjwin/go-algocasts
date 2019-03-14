package leetcode199

import (
	"fmt"
	"testing"

	"github.com/cxjwin/go-algocasts/datastructure"
)

func TestRightSideView(t *testing.T) {
	root := ds.BuildTree([]interface{}{1, 2, 3, nil, 5, nil, 4})
	res := rightSideView(root)
	fmt.Println(res)

	root = &ds.TreeNode{Val: 1}
	root.Left = &ds.TreeNode{Val: 2}
	res = rightSideView(root)
	fmt.Println(res)
}

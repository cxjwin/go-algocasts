package leetcode105

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTreeWithMap(pre []int, preStart int, preEnd int, inMap map[int]int, inStart int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := pre[preStart]
	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}
	rootIdx := inMap[rootVal]
	leftLen := rootIdx - inStart
	root.Left = buildTreeWithMap(pre, preStart+1, preStart+leftLen, inMap, inStart)
	root.Right = buildTreeWithMap(pre, preStart+leftLen+1, preEnd, inMap, rootIdx+1)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}

	return buildTreeWithMap(preorder, 0, len(preorder)-1, inMap, 0)
}

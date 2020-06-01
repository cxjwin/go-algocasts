package leetcode106

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func buildTreeWithMap(post []int, postStart int, postEnd int, inMap map[int]int, inStart int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	rootVal := post[postEnd]
	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}
	rootIdx := inMap[rootVal]
	leftLen := rootIdx - inStart
	root.Left = buildTreeWithMap(post, postStart, postStart+leftLen-1, inMap, inStart)
	root.Right = buildTreeWithMap(post, postStart+leftLen, postEnd-1, inMap, rootIdx+1)
	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}
	return buildTreeWithMap(postorder, 0, len(postorder)-1, inMap, 0)
}

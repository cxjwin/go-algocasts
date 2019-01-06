package algo

import (
	"github.com/cxjwin/go-algocasts/datastructure"
	"github.com/cxjwin/go-algocasts/utils"
	"github.com/golang-collections/collections/queue"
)

func minDepthOfBinaryTreeRecursive(tree *ds.Tree) int {
	if tree == nil {
		return 0
	}

	if tree.Left == nil && tree.Right == nil {
		return 1
	}

	if tree.Left != nil && tree.Right == nil {
		return minDepthOfBinaryTreeRecursive(tree.Left) + 1
	}

	if tree.Right != nil && tree.Left == nil {
		return minDepthOfBinaryTreeRecursive(tree.Right) + 1
	}

	return utils.IntMin(minDepthOfBinaryTreeRecursive(tree.Left), minDepthOfBinaryTreeRecursive(tree.Right)) + 1
}

func minDepthOfBinaryTreeIterative(tree *ds.Tree) int {
	if tree == nil {
		return 0
	}

	depth := 1
	queue := queue.New()
	queue.Enqueue(tree)

	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Dequeue().(*ds.Tree)
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
			depth++
		}
	}

	return depth
}

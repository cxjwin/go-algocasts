package algo

import (
	"github.com/cxjwin/go-algocasts/datastructure"
	"github.com/cxjwin/go-algocasts/utils"
	"github.com/golang-collections/collections/queue"
)

func maxDepthOfBinaryTree(root *ds.Tree) int {
	if root == nil {
		return 0
	}

	return utils.IntMax(maxDepthOfBinaryTree(root.Left), maxDepthOfBinaryTree(root.Right)) + 1
}

func maxDepthOfBinaryTreeIterative(root *ds.Tree) int {
	if root == nil {
		return 0
	}

	queue := queue.New()
	queue.Enqueue(root)
	depth := 0

	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Dequeue().(*ds.Tree)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		depth++
	}

	return depth
}

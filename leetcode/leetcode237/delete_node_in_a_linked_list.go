package leetcode237

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/delete-node-in-a-linked-list/

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next.Next = next.Next
	next.Next = nil
}

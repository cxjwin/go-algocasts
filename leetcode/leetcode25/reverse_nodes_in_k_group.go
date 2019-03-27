package leetcode25

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/

func reverse(p *ListNode, q *ListNode, k int) *ListNode {
	cur := p
	pre := q
	for i := 0; i < k; i++ {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	count := 0
	p := head
	for p != nil {
		count++
		p = p.Next
	}

	loop := count / k

	cur := head
	var pre *ListNode

	for i := 0; i < loop; i++ {
		q := cur
		for j := 0; j < k; j++ {
			q = q.Next
		}
		next := q

		n := reverse(cur, pre, k)
		if i == 0 {
			head = n
		}
		if pre != nil {
			pre.Next = n
		}

		pre = cur
		cur = next
	}

	if pre != nil {
		pre.Next = cur
	}

	return head
}

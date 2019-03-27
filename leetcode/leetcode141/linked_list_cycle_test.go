package leetcode141

import "testing"
import . "github.com/cxjwin/go-algocasts/datastructure"

func TestHasCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	res := hasCycle(head)
	if !res {
		t.Error("Input: head = [3,2,0,-4], pos = 1; Output: true")
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = head

	res = hasCycle(head)
	if !res {
		t.Error("Input: head = [1,2], pos = 0; Output: true")
	}
}

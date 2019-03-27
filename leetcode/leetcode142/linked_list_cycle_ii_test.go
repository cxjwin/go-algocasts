package leetcode142

import "testing"
import . "github.com/cxjwin/go-algocasts/datastructure"

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	res := detectCycle(head)
	if res != head.Next {
		t.Error("Input: head = [3,2,0,-4], pos = 1; Output: 1")
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = head

	res = detectCycle(head)
	if res != head {
		t.Error("Input: head = [1,2], pos = 0; Output: 0")
	}
}

package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/add-two-numbers/

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	p := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: sum % 10, Next: nil}
		p = p.Next
		carry = sum / 10
	}

	return dummy.Next
}

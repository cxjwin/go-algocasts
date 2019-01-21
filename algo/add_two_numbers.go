package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func addTwoNumbers(l1, l2 *ds.ListNode) *ds.ListNode {
	dummy := &ds.ListNode{Value: 0, Next: nil}
	p := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Value.(int)
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Value.(int)
			l2 = l2.Next
		}

		p.Next = &ds.ListNode{Value: sum % 10, Next: nil}
		p = p.Next
		carry = sum / 10
	}

	return dummy.Next
}

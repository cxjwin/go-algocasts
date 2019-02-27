package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKSortedLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	k := len(lists)
	heads := make([]*ListNode, k)

	for i := 0; i < k; i++ {
		heads[i] = lists[i]
	}

	dummy := &ListNode{Val: 0, Next: nil}
	p := dummy
	for {
		minIdx := -1
		for i := 0; i < k; i++ {
			if heads[i] != nil {
				if minIdx == -1 {
					minIdx = i
				} else {
					if heads[i].Val < heads[minIdx].Val {
						minIdx = i
					}
				}
			}
		}

		if minIdx == -1 {
			break
		} else {
			p.Next = heads[minIdx]
			p = p.Next
			heads[minIdx] = heads[minIdx].Next
		}
	}

	return dummy.Next
}

func mergeListsRecursive(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	l1 := mergeListsRecursive(lists, l, mid)
	l2 := mergeListsRecursive(lists, mid+1, r)
	return mergeTwoSortedList(l1, l2)
}

func mergeKSortedListsDivideConquer(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return mergeListsRecursive(lists, 0, len(lists)-1)
}

func mergeKSortedListsOneByOne(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	k := len(lists)
	var head *ListNode
	for i := 0; i < k; i++ {
		if head == nil {
			head = lists[i]
		} else {
			head = mergeTwoSortedList(head, lists[i])
		}
	}

	return head
}

func addToMinSlice(nodes []*ListNode, node *ListNode) []*ListNode {
	if nodes == nil {
		return []*ListNode{node}
	}

	if len(nodes) == 0 {
		return append(nodes, node)
	}

	idx := len(nodes)
	for i := 0; i < len(nodes); i++ {
		if node.Val <= nodes[i].Val {
			idx = i
			break
		}
	}

	s := append(nodes, nil /* use the zero value of the element type */)
	copy(s[idx+1:], s[idx:])
	s[idx] = node

	return s
}

func removeFristFromMinSlice(nodes []*ListNode) []*ListNode {
	if nodes == nil || len(nodes) == 0 {
		return nodes
	}

	return nodes[1:]
}

func mergeKSortedListsMinClice(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	heads := make([]*ListNode, 0)

	for _, v := range lists {
		if v != nil {
			heads = addToMinSlice(heads, v)
		}
	}

	dummy := &ListNode{Val: 0, Next: nil}
	p := dummy

	for {
		if len(heads) == 0 {
			break
		}

		min := heads[0]
		p.Next = min
		p = p.Next

		heads = removeFristFromMinSlice(heads)
		if min.Next != nil {
			heads = addToMinSlice(heads, min.Next)
		}
	}

	return dummy.Next
}

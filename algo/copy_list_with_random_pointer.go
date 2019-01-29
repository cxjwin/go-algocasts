package algo

// RandomListNode struct
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func copyListWithRandomPointer(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	m := make(map[*RandomListNode]*RandomListNode)

	copyHead := &RandomListNode{Val: head.Val}
	m[head] = copyHead

	p := head.Next
	cp := copyHead

	for p != nil {
		cp.Next = &RandomListNode{Val: p.Val}
		m[p] = cp.Next
		p = p.Next
		cp = cp.Next
	}

	p = head
	cp = copyHead
	for p != nil {
		if p.Random != nil {
			cr := m[p.Random]
			cp.Random = cr
		}

		p = p.Next
		cp = cp.Next
	}

	return copyHead
}

func copyListWithRandomPointerO1(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	// 1. double list
	p := head
	for p != nil {
		next := p.Next
		cp := &RandomListNode{Val: p.Val}
		p.Next = cp
		cp.Next = next
		p = next
	}

	// 2. double random
	p = head
	for p != nil {
		cp := p.Next
		next := cp.Next
		r := p.Random
		if r != nil {
			cp.Random = r.Next
		}
		p = next
	}

	// 3. cut pointer
	p = head
	dummy := &RandomListNode{}
	cp := dummy
	for p != nil {
		copy := p.Next

		cp.Next = copy
		cp = copy

		next := copy.Next
		p.Next = next
		p = next
	}

	return dummy.Next
}

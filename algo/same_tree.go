package algo

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
	"github.com/golang-collections/collections/stack"
)

func isSameTreeRecusive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTreeRecusive(p.Left, q.Left) &&
		isSameTreeRecusive(p.Right, q.Right)
}

func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	s := new(stack.Stack)
	s.Push(p)
	s.Push(q)

	for s.Len() != 0 {
		n1 := s.Pop()
		n2 := s.Pop()
		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		tN1 := n1.(*TreeNode)
		tN2 := n2.(*TreeNode)
		if tN1.Val != tN2.Val {
			return false
		}

		if tN1.Left != nil {
			s.Push(tN1.Left)
		}
		if tN2.Left != nil {
			s.Push(tN2.Left)
		}
		if tN1.Right != nil {
			s.Push(tN1.Right)
		}
		if tN2.Right != nil {
			s.Push(tN2.Right)
		}
	}

	return true
}

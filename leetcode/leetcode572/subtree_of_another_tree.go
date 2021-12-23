package leetcode572

import (
	"hash/fnv"
	"strconv"

	. "github.com/cxjwin/go-algocasts/datastructure"
)

// https://leetcode.com/problems/subtree-of-another-tree/

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val &&
		isSameTree(s.Left, t.Left) &&
		isSameTree(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func computeHash(root *TreeNode, m map[*TreeNode]uint32) string {
	if root == nil {
		return "#"
	}

	left := computeHash(root.Left, m)
	right := computeHash(root.Right, m)
	hashString := left + strconv.Itoa(root.Val) + right
	m[root] = hash(hashString)
	return hashString
}

func isSub(s *TreeNode, t *TreeNode, m map[*TreeNode]uint32) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	return m[s] == m[t] || isSub(s.Left, t, m) || isSub(s.Right, t, m)
}

func isSubtreeHash(s *TreeNode, t *TreeNode) bool {
	m := make(map[*TreeNode]uint32)
	computeHash(s, m)
	computeHash(t, m)
	return isSub(s, t, m)
}

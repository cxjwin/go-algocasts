package leetcode146

import (
	"fmt"
	"testing"
)

func TestMove2Head(t *testing.T) {
	head := &node{key: 1, val: 1}

	cur := head
	for i := 2; i <= 4; i++ {
		node := &node{key: i, val: i}
		cur.next = node
		node.prev = cur
		cur = node
	}

	head.prev = cur
	cur.next = head

	fmt.Println(head)
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Error("case1 : Get(1) -> 1")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Error("case2 : Get(2) -> -1")
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Error("case3 : Get(1) -> -1")
	}
	if cache.Get(3) != 3 {
		t.Error("case4 : Get(3) -> 3")
	}
	if cache.Get(4) != 4 {
		t.Error("case5 : Get(4) -> 4")
	}
}

package ds

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	nums := []int{6, 5, 4, 3, 7, 8, 9}

	h := make(IntHeap, 0)

	for _, v := range nums {
		heap.Push(&h, v)
	}

	idx := h.GetIndex(9)
	heap.Remove(&h, idx)

	for i := 0; i < len(nums)-1; i++ {
		e := h.Peek()
		fmt.Println(e)
		t := heap.Pop(&h)
		fmt.Println(t)
	}
}

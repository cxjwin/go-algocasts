package ds

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push a int value
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop a int value
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek a int value
func (h *IntHeap) Peek() interface{} {
	return (*h)[0]
}

// GetIndex of a int value
func (h *IntHeap) GetIndex(v interface{}) int {
	for i, x := range *h {
		if x == v {
			return i
		}
	}
	return -1
}

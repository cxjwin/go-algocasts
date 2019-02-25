package leetcode218

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/the-skyline-problem/

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func getSkyline(buildings [][]int) [][]int {
	if buildings == nil || len(buildings) == 0 {
		return nil
	}

	height := make([][]int, len(buildings)*2)
	for i, v := range buildings {
		height[i*2] = []int{v[0], -v[2]}
		height[i*2+1] = []int{v[1], v[2]}
	}

	sort.Slice(height, func(i, j int) bool {
		a := height[i]
		b := height[j]
		if a[0] == b[0] {
			return a[1]-b[1] < 0
		}

		return a[0]-b[0] < 0
	})

	result := make([][]int, 0)

	hp := make(IntHeap, 0)
	heap.Push(&hp, 0)

	preHeight := 0

	for _, h := range height {
		if h[1] < 0 {
			heap.Push(&hp, -h[1])
		} else {
			idx := hp.GetIndex(h[1])
			if idx != -1 {
				heap.Remove(&hp, idx)
			}
		}

		curHeight := hp.Peek()
		if curHeight != preHeight {
			result = append(result, []int{h[0], curHeight.(int)})
			preHeight = curHeight.(int)
		}
	}

	return result
}

func getSkylineWithMap(buildings [][]int) [][]int {
	if buildings == nil || len(buildings) == 0 {
		return nil
	}

	height := make([][]int, len(buildings)*2)
	for i, v := range buildings {
		height[i*2] = []int{v[0], -v[2]}
		height[i*2+1] = []int{v[1], v[2]}
	}

	sort.Slice(height, func(i, j int) bool {
		a := height[i]
		b := height[j]
		if a[0] == b[0] {
			return a[1]-b[1] < 0
		}

		return a[0]-b[0] < 0
	})

	result := make([][]int, 0)

	m := make(map[int]int)
	m[0] = 1
	preHeight := 0

	for _, h := range height {
		if h[1] < 0 {
			k := -h[1]
			v, ok := m[k]
			if ok {
				m[k] = v + 1
			} else {
				m[k] = 1
			}
		} else {
			k := h[1]
			v, ok := m[k]
			if ok && v == 1 {
				delete(m, k)
			} else {
				m[k] = v - 1
			}
		}

		// get largest key, O(n)
		largestKey := 0
		for k := range m {
			if k > largestKey {
				largestKey = k
			}
		}

		if preHeight != largestKey {
			result = append(result, []int{h[0], largestKey})
			preHeight = largestKey
		}
	}

	return result
}

package algo

import "sort"

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

// A KthLargest is stauct to find the kth largest element in a stream
type KthLargest struct {
	minHeap []int
	k       int
}

// NewKthLargest a KthLargest struct
func NewKthLargest(k int, nums []int) KthLargest {
	kl := KthLargest{}
	kl.minHeap = make([]int, 0)
	kl.k = k
	for _, v := range nums {
		kl.Add(v)
	}

	return kl
}

// Add a Value
func (kl *KthLargest) Add(val int) int {
	if len(kl.minHeap) < kl.k {
		kl.minHeap = append(kl.minHeap, val)
		sort.Ints(kl.minHeap)
	} else {
		if val > kl.minHeap[0] {
			kl.minHeap[0] = val
			sort.Ints(kl.minHeap)
		}
	}

	return kl.minHeap[0]
}

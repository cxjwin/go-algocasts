package leetcode384

import "math/rand"

// Solution for leetcode 384
type Solution struct {
	orig []int
	nums []int
}

// Constructor a Solution struct
func Constructor(nums []int) Solution {
	s := Solution{}
	s.orig = nums
	s.nums = make([]int, len(nums))
	copy(s.nums, nums)
	return s
}

// Reset - Resets the array to its original configuration and return it.
func (s *Solution) Reset() []int {
	copy(s.nums, s.orig)
	return s.nums
}

// Shuffle - Returns a random shuffling of the array.
func (s *Solution) Shuffle() []int {
	for i := len(s.nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		t := s.nums[i]
		s.nums[i] = s.nums[j]
		s.nums[j] = t
	}
	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

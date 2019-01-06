package algo

import "github.com/cxjwin/go-algocasts/utils"

func medianOfTwoSortedArrays(nums1 []int, nums2 []int) float32 {
	totalLen := len(nums1) + len(nums2)
	if (totalLen & 1) == 1 {
		return float32(findKthSmallestNum(nums1, nums2, totalLen/2+1))
	}

	a := findKthSmallestNum(nums1, nums2, totalLen/2)
	b := findKthSmallestNum(nums1, nums2, totalLen/2+1)
	return float32(a+b) / 2.0
}

func findKthSmallestNum(nums1 []int, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	base1 := 0
	base2 := 0

	var res int
	for true {
		if len1 == 0 {
			res = nums2[base2+k-1]
			break
		}
		if len2 == 0 {
			res = nums1[base1+k-1]
			break
		}
		if k == 1 {
			res = utils.IntMin(nums1[base1], nums2[base2])
			break
		}

		i := utils.IntMin(k/2, len1)
		j := utils.IntMin(k-i, len2)
		a := nums1[base1+i-1]
		b := nums2[base2+j-1]

		if i+j == k && a == b {
			res = a
			break
		}

		if a <= b {
			base1 += i
			len1 -= i
			k -= i
		}
		if a >= b {
			base2 += j
			len2 -= j
			k -= j
		}
	}

	return res
}

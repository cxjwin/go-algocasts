package sort

// https://www.geeksforgeeks.org/bucket-sort-2/

import "sort"

func bucketSort(nums []float32, n int) {
	// 1. Create n empty buckets
	buckets := make([][]float32, n)
	for i := 0; i < n; i++ {
		buckets[i] = make([]float32, 0)
	}

	// 2. Put array elements in different buckets
	for _, v := range nums {
		idx := int(v * float32(n))
		buckets[idx] = append(buckets[idx], v)
	}

	// 3. Sort individual buckets
	for i := 0; i < n; i++ {
		temp := buckets[i]
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
	}

	// 4. Concatenate all buckets into nums
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			nums[index] = buckets[i][j]
			index++
		}
	}
}

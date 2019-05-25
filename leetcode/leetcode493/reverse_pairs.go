package leetcode493

func reversePairsBruteForce(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j]*2 {
				count++
			}
		}
	}

	return count
}

func merge(nums []int, low int, mid int, high int) {
	m, n := mid-low+1, high-mid
	a1 := make([]int, m)
	for i := 0; i < m; i++ {
		a1[i] = nums[low+i]
	}
	a2 := make([]int, n)
	for i := 0; i < n; i++ {
		a2[i] = nums[mid+1+i]
	}

	i, j, k := 0, 0, low
	for i < m && j < n {
		if a1[i] < a2[j] {
			nums[k] = a1[i]
			i++
		} else {
			nums[k] = a2[j]
			j++
		}
		k++
	}
	for i < m {
		nums[k] = a1[i]
		i++
		k++
	}
	for j < n {
		nums[k] = a2[j]
		j++
		k++
	}
}

func count(nums []int, low int, mid int, high int) int {
	cnt := 0
	i, j := low, mid+1
	for i <= mid && j <= high {
		if nums[i] <= nums[j]*2 {
			i++
		} else {
			cnt += (mid - i + 1)
			j++
		}
	}
	return cnt
}

func sortAndCount(nums []int, low int, high int) int {
	if low >= high {
		return 0
	}

	cnt := 0
	mid := low + (high-low)/2
	cnt += sortAndCount(nums, low, mid)
	cnt += sortAndCount(nums, mid+1, high)
	cnt += count(nums, low, mid, high)
	merge(nums, low, mid, high)

	return cnt
}

func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	return sortAndCount(nums, 0, len(nums)-1)
}

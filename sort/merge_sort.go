package sort

// https://www.geeksforgeeks.org/merge-sort/

func merge(nums []int, l int, m int, r int) {
	n1 := m - l + 1
	leftArr := make([]int, n1)
	for i := 0; i < n1; i++ {
		leftArr[i] = nums[l+i]
	}
	n2 := r - m
	rightArr := make([]int, n2)
	for i := 0; i < n2; i++ {
		rightArr[i] = nums[m+1+i]
	}

	i := 0
	j := 0
	k := l

	for i < n1 && j < n2 {
		if leftArr[i] < rightArr[j] {
			nums[k] = leftArr[i]
			i++
			k++
		} else {
			nums[k] = rightArr[j]
			j++
			k++
		}
	}

	for i < n1 {
		nums[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		nums[k] = rightArr[j]
		j++
		k++
	}
}

func mergeWithGuard(nums []int, l int, m int, r int) {
	guard := 1<<31 - 1
	n1 := m - l + 1
	leftArr := make([]int, n1+1)
	for i := 0; i < n1; i++ {
		leftArr[i] = nums[l+i]
	}
	leftArr[n1] = guard
	n2 := r - m
	rightArr := make([]int, n2+1)
	for i := 0; i < n2; i++ {
		rightArr[i] = nums[m+i+1]
	}
	rightArr[n2] = guard

	i, j := 0, 0
	for k := 0; k < r-l; k++ {
		if leftArr[i] <= rightArr[j] {
			nums[k] = leftArr[i]
			i++
		} else {
			nums[k] = rightArr[j]
			j++
		}
	}
}

func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	// Same as (l+r)/2, but avoids overflow for
	// large l and h
	m := l + (r-l)/2

	// Sort first and second halves
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)

	// merge
	merge(nums, l, m, r)
}

func mergeSortWithGuard(nums []int, l int, r int) {
	if l >= r {
		return
	}

	// Same as (l+r)/2, but avoids overflow for
	// large l and h
	m := l + (r-l)/2

	// Sort first and second halves
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)

	// merge
	mergeWithGuard(nums, l, m, r)
}

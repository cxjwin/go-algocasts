package sort

// https://www.geeksforgeeks.org/insertion-sort/

func insertionSort(arr []int, num int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		k := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > k {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = k
	}
}

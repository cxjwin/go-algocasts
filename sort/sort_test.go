package sort

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func getFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func TestSortFunc(t *testing.T) {
	type testFunc func(nums []int)

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
		res := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 12}

		copyNums := make([]int, len(nums))
		copy(copyNums, nums)

		f(copyNums)

		if !reflect.DeepEqual(copyNums, res) {
			fmt.Println(copyNums)
			t.Error(getFunctionName(f))
		}
	}

	// bubble sort
	// testBody(bubbleSort, t)
	// testBody(bubbleSort2, t)
	// testBody(cocktailSort, t)

	// selection sort
	// testBody(selectionSort, t)

	// quick sort
	// testBody(quickSort1, t)
	// testBody(quickSort2, t)
	// testBody(quickSort3, t)

	// testBody(countingSort, t)

	testBody(radixSort, t)
}

func TestBucketSort(t *testing.T) {
	nums := []float32{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	bucketSort(nums, len(nums))

	res := []float32{0.1234, 0.3434, 0.565, 0.656, 0.665, 0.897}

	if !reflect.DeepEqual(nums, res) {
		fmt.Println(nums)
		t.Error("bucketSort")
	}
}

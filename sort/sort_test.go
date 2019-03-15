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

func TestBubbleSort(t *testing.T) {
	type testFunc func(nums []int)

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
		res := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 12}

		f(nums)

		if !reflect.DeepEqual(nums, res) {
			fmt.Println(nums)
			t.Error(getFunctionName(f))
		}
	}

	testBody(bubbleSort, t)
}

func TestInsertitonSort(t *testing.T) {
	nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
	res := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 12}

	insertionSort(nums)

	if !reflect.DeepEqual(nums, res) {
		fmt.Println(nums)
		t.Error(getFunctionName(insertionSort))
	}
}

func TestQuickSort(t *testing.T) {
	type testFunc func(nums []int)

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
		res := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 12}

		f(nums)

		if !reflect.DeepEqual(nums, res) {
			fmt.Println(nums)
			t.Error(getFunctionName(f))
		}
	}

	testBody(quickSort1, t)
	testBody(quickSort3, t)
}

func TestMergeSort(t *testing.T) {
	type testFunc func([]int, int, int)

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
		res := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 12}

		f(nums, 0, len(nums)-1)

		if !reflect.DeepEqual(nums, res) {
			fmt.Println(nums)
			t.Error("mergeSort")
		}
	}

	testBody(mergeSort, t)
	testBody(mergeSortWithGuard, t)
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

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
			t.Error(getFunctionName(bubbleSort2))
		}
	}

	// testBody(bubbleSort, t)
	// testBody(bubbleSort2, t)
	// testBody(cocktailSort, t)

	// testBody(selectionSort, t)

	// testBody(quickSort1, t)
	// testBody(quickSort2, t)
	testBody(quickSort3, t)
}

func TestPartion(t *testing.T) {
	nums := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
	i := partion2(nums, 0, len(nums)-1)
	fmt.Println(i)

	nums2 := []int{1, 4, 5, 6, 2, 9, 8, 7, 3, 5, 2, 12}
	right := len(nums2) - 1
	j := partion3(nums2, 0, right, right)
	fmt.Println(j)
}

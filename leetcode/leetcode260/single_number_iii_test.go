package leetcode260

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSingleNumberIII(t *testing.T) {
	nums := []int{2, 4, 2, 1, 4, 8}
	res := singleNumber(nums)
	fmt.Println(res)
	sort.Ints(res)
	if !reflect.DeepEqual(res, []int{1, 8}) {
		t.Error("2, 4, 2, 1, 4, 8 -> 1, 8")
	}
}

func TestSingleNumberIII2(t *testing.T) {
	nums := []int{2, 4, 2, 1, 4, 8}
	res := singleNumber2(nums)
	fmt.Println(res)
	sort.Ints(res)
	if !reflect.DeepEqual(res, []int{1, 8}) {
		t.Error("2, 4, 2, 1, 4, 8 -> 1, 8")
	}
}

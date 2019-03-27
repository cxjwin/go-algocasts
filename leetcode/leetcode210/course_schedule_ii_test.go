package leetcode210

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	case1 := [][]int{
		{1, 0},
		{3, 0},
		{3, 1},
		{2, 1},
		{2, 3},
		{4, 2},
		{4, 3},
	}
	if !reflect.DeepEqual(findOrder(5, case1), []int{0, 1, 3, 2, 4}) {
		t.Error("case1 => 0, 1, 3, 2, 4")
	}

	case2 := [][]int{
		{1, 0},
	}
	if !reflect.DeepEqual(findOrder(2, case2), []int{0, 1}) {
		t.Error("case2 => 0, 1")
	}

	case3 := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	res := findOrder(4, case3)
	if !reflect.DeepEqual(res, []int{0, 1, 2, 3}) &&
		!reflect.DeepEqual(res, []int{0, 2, 1, 3}) {
		t.Error("case3 => 0, 1, 2, 3 or 0, 2, 1, 3")
	}
}

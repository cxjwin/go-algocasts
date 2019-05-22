package leetcode54

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := spiralOrder(matrix)
	if !reflect.DeepEqual(res, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}) {
		t.Error("=> 1,2,3,6,9,8,7,4,5")
	}

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res = spiralOrder(matrix)
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}) {
		t.Error("=> 1,2,3,4,8,12,11,10,9,5,6,7")
	}
}

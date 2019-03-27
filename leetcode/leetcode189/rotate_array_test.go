package leetcode189

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{0, 1, 2, 4, 8}
	rotate(nums, 3)
	if !reflect.DeepEqual(nums, []int{2, 4, 8, 0, 1}) {
		t.Error("0, 1, 2, 4, 8 -- (3) --> 2, 4, 8, 0, 1")
	}
}

func TestMove2Head(t *testing.T) {
	nums := []int{0, 1, 2, 4, 8}
	move2Head(nums, 0, 2)
	fmt.Println(nums)
}

func TestRotate2(t *testing.T) {
	nums := []int{0, 1, 2, 4, 8}
	rotate2(nums, 3)
	if !reflect.DeepEqual(nums, []int{2, 4, 8, 0, 1}) {
		t.Error("0, 1, 2, 4, 8 -- (3) --> 2, 4, 8, 0, 1")
	}
}

func TestRotate3(t *testing.T) {
	nums := []int{0, 1, 2, 4, 8}
	rotate3(nums, 3)
	if !reflect.DeepEqual(nums, []int{2, 4, 8, 0, 1}) {
		t.Error("0, 1, 2, 4, 8 -- (3) --> 2, 4, 8, 0, 1")
	}
}

package leetcode338

import (
	"reflect"
	"testing"
)

func TestCountBitsOfNumber(t *testing.T) {
	if countBitsOfNumber(1) != 1 {
		t.Error("1 => 1")
	}
	if countBitsOfNumber(3) != 2 {
		t.Error("3 => 2")
	}
}

func TestCountBits(t *testing.T) {
	type testFunc func(int) []int

	testBody := func(f testFunc, t *testing.T) {
		if !reflect.DeepEqual(f(2), []int{0, 1, 1}) {
			t.Error("2 => [0, 1, 1]")
		}
		if !reflect.DeepEqual(f(5), []int{0, 1, 1, 2, 1, 2}) {
			t.Error("5 => [0, 1, 1, 2, 1, 2]")
		}
	}

	testBody(countBits, t)
	testBody(countBitsDP, t)
}

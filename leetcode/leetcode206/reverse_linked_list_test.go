package leetcode206

import (
	"reflect"
	"testing"

	. "github.com/cxjwin/go-algocasts/datastructure"
)

func TestReverseLinkedList(t *testing.T) {
	type testFunc func(*ListNode) *ListNode

	testBody := func(f testFunc, t *testing.T) {
		head := Array2List([]int{1, 2, 3, 4, 5})
		n := f(head)
		res := List2Array(n)
		if !reflect.DeepEqual(res, []int{5, 4, 3, 2, 1}) {
			t.Error("Output: 5->4->3->2->1->nil")
		}
	}

	// testBody(reverseList, t)
	testBody(reverseListIterative, t)
}

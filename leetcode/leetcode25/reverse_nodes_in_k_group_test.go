package leetcode25

import (
	"reflect"
	"testing"

	. "github.com/cxjwin/go-algocasts/datastructure"
)

func TestReverse(t *testing.T) {
	head := Array2List([]int{1, 2, 3, 4, 5})
	head = reverse(head, nil, 3)
	if !reflect.DeepEqual(List2Array(head), []int{3, 2, 1}) {
		t.Error("Input: 1->2->3->4->5->NULL, k = 3, Output: 3->2->1->NULL")
	}
}

func TestReverseKGroup(t *testing.T) {
	var head *ListNode
	head = Array2List([]int{1, 2, 3, 4, 5})
	head = reverseKGroup(head, 3)
	if !reflect.DeepEqual(List2Array(head), []int{3, 2, 1, 4, 5}) {
		t.Error("Input: 1->2->3->4->5->NULL, k = 3, Output: 3->2->1->4->5->NULL")
	}
	head = Array2List([]int{1, 2, 3, 4, 5})
	head = reverseKGroup(head, 2)
	if !reflect.DeepEqual(List2Array(head), []int{2, 1, 4, 3, 5}) {
		t.Error("Input: 1->2->3->4->5->NULL, k = 3, Output: 2->1->4->3->5->NULL")
	}
	head = Array2List([]int{1, 2, 3, 4})
	head = reverseKGroup(head, 2)
	if !reflect.DeepEqual(List2Array(head), []int{2, 1, 4, 3}) {
		t.Error("Input: 1->2->3->4->NULL, k = 3, Output: 2->1->4->3->NULL")
	}
}

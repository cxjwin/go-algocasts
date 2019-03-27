package leetcode24

import (
	"reflect"
	"testing"

	. "github.com/cxjwin/go-algocasts/datastructure"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/

func TestSwapPairs(t *testing.T) {
	head := Array2List([]int{1, 2, 3, 4})
	head = swapPairs(head)
	if !reflect.DeepEqual(List2Array(head), []int{2, 1, 4, 3}) {
		t.Error("Given 1->2->3->4, you should return the list as 2->1->4->3.")
	}

	head = Array2List([]int{1, 2, 3, 4, 5})
	head = swapPairs(head)
	if !reflect.DeepEqual(List2Array(head), []int{2, 1, 4, 3, 5}) {
		t.Error("Given 1->2->3->4->5, you should return the list as 2->1->4->3->5.")
	}
}

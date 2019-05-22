package lintcode1307

import "testing"

func TestVerifyPreorderDivideConquer(t *testing.T) {
	if verifyPreorderDivideConquer([]int{4, 1, 0, 2, 8}) != true {
		t.Error("[4, 1, 0, 2, 8] is preorder")
	}

	if verifyPreorderDivideConquer([]int{4, 1, 0, 8, 2}) == true {
		t.Error("[4, 1, 0, 2, 8] is not preorder")
	}
}

func TestVerifyPreorderStack(t *testing.T) {
	if verifyPreorderStack([]int{4, 1, 0, 2, 8}) != true {
		t.Error("[4, 1, 0, 2, 8] is preorder")
	}

	if verifyPreorderStack([]int{4, 1, 0, 8, 2}) == true {
		t.Error("[4, 1, 0, 2, 8] is not preorder")
	}
}

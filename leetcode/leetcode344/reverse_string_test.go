package leetcode344

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(str)
	reverseString(str)
	fmt.Println(str)

}

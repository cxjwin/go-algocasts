package utils

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	// TODO
	num := 12341
	m := 64

	v1 := num % m
	v2 := num & (m - 1)

	fmt.Println(v1, v2)
}

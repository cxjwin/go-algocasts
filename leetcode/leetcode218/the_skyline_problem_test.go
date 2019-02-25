package leetcode218

import (
	"fmt"
	"testing"
)

func TestSkyline(t *testing.T) {
	// [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ]
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}

	// [ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]
	res := getSkyline(buildings)

	fmt.Println(res)

	if len(res) != 7 {
		t.Error("7 points")
	}

	if res[0][0] != 2 || res[0][1] != 10 {
		t.Error("1 : [2, 10]")
	}
	if res[1][0] != 3 || res[1][1] != 15 {
		t.Error("1 : [3, 15]")
	}
	if res[2][0] != 7 || res[2][1] != 12 {
		t.Error("1 : [7, 12]")
	}
	if res[3][0] != 12 || res[3][1] != 0 {
		t.Error("1 : [12, 0]")
	}
	if res[4][0] != 15 || res[4][1] != 10 {
		t.Error("1 : [15, 10]")
	}
	if res[5][0] != 20 || res[5][1] != 8 {
		t.Error("1 : [20, 8]")
	}
	if res[6][0] != 24 || res[6][1] != 0 {
		t.Error("1 : [24, 0]")
	}
}

package leetcode56

import "testing"

func TestMergeIntervals(t *testing.T) {
	// Input: [[1,3],[2,6],[8,10],[15,18]]
	i1 := Interval{Start: 1, End: 3}
	i2 := Interval{Start: 2, End: 6}
	i3 := Interval{Start: 8, End: 10}
	i4 := Interval{Start: 15, End: 18}
	res := merge([]Interval{i1, i2, i3, i4})

	if len(res) != 3 {
		t.Error("3 merged instervals")
	}

	// Output: [[1,6],[8,10],[15,18]]
	r1 := res[0]
	if r1.Start != 1 || r1.End != 6 {
		t.Error("[1,6]")
	}
	r2 := res[1]
	if r2.Start != 8 || r2.End != 10 {
		t.Error("[8,10]")
	}
	r3 := res[2]
	if r3.Start != 15 || r3.End != 18 {
		t.Error("[15,18]")
	}
}

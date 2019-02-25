package leetcode56

import (
	"sort"
)

// https://leetcode.com/problems/merge-intervals/

//Interval - Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}

	// sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start-intervals[j].Start < 0
	})

	res := make([]Interval, 1)
	res[0] = intervals[0]
	n := len(intervals)

	for i := 1; i < n; i++ {
		lastIdx := len(res) - 1
		interval := intervals[i]

		if interval.Start > res[lastIdx].End {
			res = append(res, interval)
		} else if interval.End > res[lastIdx].End {
			res[lastIdx].End = interval.End
		} else {
			// in interval, do nothing
		}
	}

	return res
}

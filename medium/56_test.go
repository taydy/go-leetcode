package medium

import "testing"

func TestMergeInterval(t *testing.T) {
	intervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	t.Log(MergeInterval(intervals))
}

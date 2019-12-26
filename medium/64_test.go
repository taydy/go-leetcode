package medium

import "testing"

func TestMinPathSum(t *testing.T) {
	paths := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	if actual := MinPathSum(paths); actual != 7 {
		t.Logf("except %d, actual %d", 7, actual)
		t.FailNow()
	}
}

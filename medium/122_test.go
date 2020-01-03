package medium

import "testing"

func TestMaxProfit2(t *testing.T) {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{1, 2, 3, 4, 5},
	}
	excepts := []int{
		7, 0, 4,
	}

	for i, v := range inputs {
		result := MaxProfit2(v)
		if result != excepts[i] {
			t.Logf("except %d, actual %d", excepts[i], result)
			t.FailNow()
		}
	}
}

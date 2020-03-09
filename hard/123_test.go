package hard

import "testing"

func TestMaxProfitIII(t *testing.T) {
	inputs := [][]int{
		{3, 3, 5, 0, 0, 3, 1, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}
	excepts := []int{
		6, 4, 0,
	}

	for i, v := range inputs {
		result := MaxProfitIII(v)
		if result != excepts[i] {
			t.Logf("except %d, actual %d", excepts[i], result)
			t.FailNow()
		}
	}
}

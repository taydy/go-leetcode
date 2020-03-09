package easy

import "testing"

func TestMaxProfit(t *testing.T) {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}
	excepts := []int{
		5, 0,
	}

	for i, v := range inputs {
		result := MaxProfit(v)
		if result != excepts[i] {
			t.Logf("except %d, actual %d", excepts[i], result)
			t.FailNow()
		}
	}

}

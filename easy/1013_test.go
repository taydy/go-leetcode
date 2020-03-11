package easy

import "testing"

func TestCanThreePartsEqualSum(t *testing.T) {
	inputs := [][]int{
		{1, -1, 1, -1},
		{12, -4, 16, -5, 9, -3, 3, 8, 0},
		{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
		{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
		{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
		{10, -7, 13, -14, 30, 16, -3, -16, -27, 27, 19},
	}
	excepts := []bool{
		false, true, true, false, true, true,
	}
	for i, v := range inputs {
		if result := CanThreePartsEqualSum(v); result != excepts[i] {
			t.Logf("input %v, except %v, but got wrong answer %v", v, excepts[i], result)
			t.FailNow()
		}
	}
}

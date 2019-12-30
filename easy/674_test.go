package easy

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 4, 7},
		{2, 2, 2, 2, 2},
	}
	excepts := []int{3, 1}

	for i, num := range nums {
		result := FindLengthOfLCIS(num)
		if result != excepts[i] {
			t.Logf("except %d, actual %d", excepts[i], result)
			t.FailNow()
		}
	}
}

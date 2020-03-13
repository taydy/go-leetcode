package easy

import "testing"

func TestMajorityElement(t *testing.T) {
	inputs := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}
	excepts := []int{
		3, 2,
	}
	for i, v := range inputs {
		if result := MajorityElement(v); result != excepts[i] {
			t.Logf("input %v, except %d. but got wrong answer %d", v, excepts[i], result)
			t.FailNow()
		}
	}
}

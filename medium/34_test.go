package medium

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	result := SearchRange(nums, 8)
	if result[0] != 3 || result[1] != 4 {
		t.Logf("except %#v, actual %#v", []int{3, 4}, result)
		t.FailNow()
	}

	result = SearchRange(nums, 6)
	if result[0] != -1 || result[1] != -1 {
		t.Logf("except %#v, actual %#v", []int{-1, -1}, result)
		t.FailNow()
	}

	nums = []int{1}
	result = SearchRange(nums, 1)
	if result[0] != 0 || result[1] != 0 {
		t.Logf("except %#v, actual %#v", []int{3, 4}, result)
		t.FailNow()
	}
}

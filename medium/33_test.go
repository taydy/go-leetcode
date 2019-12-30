package medium

import "testing"

func TestSearch1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	result := Search1(nums, 0)
	if result != 4 {
		t.Logf("except %d, actual %d", 4, result)
		t.FailNow()
	}

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	result = Search1(nums, 3)
	if result != -1 {
		t.Logf("except %d, actual %d", -1, result)
		t.FailNow()
	}

	nums = []int{1, 3}
	result = Search1(nums, 3)
	if result != 1 {
		t.Logf("except %d, actual %d", 1, result)
		t.FailNow()
	}

	nums = []int{8, 9, 2, 4, 3}
	result = Search1(nums, 9)
	if result != 1 {
		t.Logf("except %d, actual %d", 9, result)
		t.FailNow()
	}

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	result = Search1(nums, 0)
	if result != 3 {
		t.Logf("except %d, actual %d", 3, result)
		t.FailNow()
	}

	nums = []int{1, 1}
	result = Search1(nums, 0)
	if result != -1 {
		t.Logf("except %d, actual %d", -1, result)
		t.FailNow()
	}

	nums = []int{1, 3, 1, 1, 1}
	result = Search1(nums, 3)
	if result != 1 {
		t.Logf("except %d, actual %d", 1, result)
		t.FailNow()
	}
}

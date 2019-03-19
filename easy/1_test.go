package easy

import "testing"

func TestTwoSum(t *testing.T)  {
	nums := []int{2, 7, 11, 15}
	total := 13
	if !equal(twoSum(nums, total), []int{0, 2}) {
		t.FailNow()
	}

	total = 26
	if !equal(twoSum(nums, total), []int{2, 3}) {
		t.FailNow()
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

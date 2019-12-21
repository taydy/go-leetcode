package medium

import "testing"

func TestFindKthLargest(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2

	result := FindKthLargest(arr, k)
	if result != 5 {
		t.Logf("except %d, actual %d", 5, result)
		t.FailNow()
	}
}

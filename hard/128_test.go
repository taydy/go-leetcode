package hard

import "testing"

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	if result := LongestConsecutive(nums); result != 4 {
		t.Logf("except %d, actual %d", 4, result)
		t.FailNow()
	}
}

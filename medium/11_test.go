package medium

import "testing"

func TestMaxArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	except := 49
	if result := MaxArea(input); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}

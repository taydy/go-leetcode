package medium

import "testing"

func TestLengthOfLIS(t *testing.T) {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	except := 4
	if result := LengthOfLIS(input); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}

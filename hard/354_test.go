package hard

import "testing"

func TestMaxEnvelopes(t *testing.T) {
	envelopes := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}
	except := 3
	if result := MaxEnvelopes(envelopes); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}

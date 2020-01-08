package easy

import "testing"

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	except := 6
	actual := MaxSubArray(arr)

	if actual != except {
		t.Logf("except %d, actual %d", except, actual)
		t.FailNow()
	}
}

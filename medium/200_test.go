package medium

import "testing"

func TestNumIslands(t *testing.T) {
	actual := NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
	})
	if actual != 1 {
		t.Logf("except %d, actual %d", 1, actual)
		t.FailNow()
	}
}

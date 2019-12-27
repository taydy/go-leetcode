package medium

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	result := MaximalSquare(matrix)
	if result != 4 {
		t.Logf("except %d, actual %d", 4, result)
		t.FailNow()
	}
}

package medium

import "testing"

func TestCountSquares(t *testing.T) {
	matrix := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}
	result := CountSquares(matrix)
	if result != 7 {
		t.Logf("except %d, actual %d", 7, result)
		t.FailNow()
	}
}

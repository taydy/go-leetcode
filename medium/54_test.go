package medium

import "testing"

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Logf("%+v", SpiralOrder(matrix))
}

package medium

import "testing"

func TestMinimumTotal(t *testing.T) {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	t.Log(MinimumTotal(arr))
}

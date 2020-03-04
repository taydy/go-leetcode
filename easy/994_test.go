package easy

import "testing"

func TestOrangesRotting(t *testing.T) {
	grids := [][][]int{
		{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		},
		{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		},
		{
			{0, 2},
		},
	}
	excepts := []int{
		4, -1, 0,
	}
	for i, v := range grids {
		if result := OrangesRotting(v); result != excepts[i] {
			t.Logf("input grid %v, except minute %d, but got wrong answer %d", v, excepts[i], result)
			t.FailNow()
		}
	}
}

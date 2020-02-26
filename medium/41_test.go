package medium

import "testing"

func TestNextPermutation(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 4, 3, 2},
	}
	excepts := [][]int{
		{1, 2, 4, 3},
		{1, 2, 3, 4},
		{2, 1, 3, 4},
	}
	for i, v := range nums {
		temp := make([]int, len(v))
		copy(temp, v)
		NextPermutation(v)
		if !Equal(v, excepts[i]) {
			t.Logf("nums %v, the next except permutation id %v, but get a wrong result %v", temp, excepts[i], v)
		}
	}
}

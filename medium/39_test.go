package medium

import "testing"

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	t.Logf("%v", CombinationSum(candidates, target))

	candidates = []int{2, 3, 5}
	target = 8
	t.Logf("%v", CombinationSum(candidates, target))
}

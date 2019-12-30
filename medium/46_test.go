package medium

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	result := Permute(nums)
	t.Logf("%+v", result)
}

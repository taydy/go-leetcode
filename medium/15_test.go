package medium

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Logf("%+v", ThreeSum(nums))
}

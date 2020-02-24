package medium

import "testing"

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Logf("%v", DailyTemperatures(temperatures))
}

package medium

import "testing"

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	except := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if result := DailyTemperatures(temperatures); !Equal(result, except) {
		t.Logf("input temperatures %v, except output %v, but actual %v", temperatures, except, result)
		t.FailNow()
	}
}

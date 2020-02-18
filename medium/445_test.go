package medium

import (
	"testing"
)

func TestAddTwoNumbers2(t *testing.T) {
	l1 := []int{7, 2, 4, 3}
	l2 := []int{5, 6, 4}
	result := AddTwoNumbers2(ArrayToList(l1), ArrayToList(l2))
	if !Equal(ListToArray(result), []int{7, 8, 0, 7}) {
		t.Logf("err output %v", ListToArray(result))
		t.FailNow()
	}
}

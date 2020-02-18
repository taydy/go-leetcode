package medium

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	result := addTwoNumbers(ArrayToList(l1), ArrayToList(l2))
	if !Equal(ListToArray(result), []int{7, 0, 8}) {
		t.FailNow()
	}

	l1 = []int{1, 8}
	l2 = []int{}
	result = addTwoNumbers(ArrayToList(l1), ArrayToList(l2))
	if !Equal(ListToArray(result), []int{1, 8}) {
		t.FailNow()
	}

	l1 = []int{5}
	l2 = []int{5}
	result = addTwoNumbers(ArrayToList(l1), ArrayToList(l2))
	if !Equal(ListToArray(result), []int{0, 1}) {
		t.FailNow()
	}
}

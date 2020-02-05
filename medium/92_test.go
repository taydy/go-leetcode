package medium

import "testing"

func TestReverseBetween(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	result := ReverseBetween(arrayToList(l1), 2, 4)
	t.Log(ListToArray(result))
}

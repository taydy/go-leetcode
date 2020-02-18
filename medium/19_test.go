package medium

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1},
	}
	removeNths := []int{2, 1}
	excepts := [][]int{
		{1, 2, 3, 5},
		{},
	}

	for i, v := range inputs {
		if result := ListToArray(RemoveNthFromEnd(ArrayToList(v), removeNths[i])); !Equal(excepts[i], result) {
			t.Logf("input %#v, removeNth %d, except %#v, actual %#v", v, removeNths[i], excepts[i], result)
			t.FailNow()
		}
	}
}

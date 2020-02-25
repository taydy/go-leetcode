package hard

import (
	"github.com/taydy/go-leetcode/medium"
	"github.com/taydy/go-leetcode/structure"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	inputs := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	except := []int{1, 1, 2, 3, 4, 4, 5, 6}
	if result := medium.ListToArray(MergeKLists(getLists(inputs))); !medium.Equal(result, except) {
		t.FailNow()
	}

}

func getLists(inputs [][]int) []*structure.ListNode {
	lists := make([]*structure.ListNode, 0)
	for _, v := range inputs {
		lists = append(lists, medium.ArrayToList(v))
	}
	return lists
}

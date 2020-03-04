package easy

import (
	"github.com/taydy/go-leetcode/structure"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 4,
			},
		},
	}
	l2 := &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 3,
			Next: &structure.ListNode{
				Val: 4,
			},
		},
	}
	result := MergeTwoLists(l1, l2)
	t.Log(result)
}

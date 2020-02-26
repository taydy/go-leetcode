package medium

import (
	"fmt"
	"github.com/taydy/go-leetcode/structure"
	"strings"
	"testing"
)

func TestLinkedSortList(t *testing.T) {
	l1 := &structure.ListNode{
		Val: 4,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 1,
				Next: &structure.ListNode{
					Val: 3,
				},
			},
		},
	}
	t.Logf("%s", print(LinkedSortList(l1)))

	l2 := &structure.ListNode{
		Val: -1,
		Next: &structure.ListNode{
			Val: 5,
			Next: &structure.ListNode{
				Val: 3,
				Next: &structure.ListNode{
					Val: 4,
					Next: &structure.ListNode{
						Val: 0,
					},
				},
			},
		},
	}
	t.Logf("%s", print(LinkedSortList(l2)))
}

func print(root *structure.ListNode) string {
	var str strings.Builder
	for root != nil {
		str.WriteString(fmt.Sprintf(" %d ", root.Val))
		root = root.Next
	}
	return str.String()
}

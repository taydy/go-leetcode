package medium

import (
	"fmt"
	"strings"
	"testing"
)

func TestLinkedSortList(t *testing.T) {
	l1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	t.Logf("%s", print(LinkedSortList(l1)))

	l2 := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}
	t.Logf("%s", print(LinkedSortList(l2)))
}

func print(root *ListNode) string {
	var str strings.Builder
	for root != nil {
		str.WriteString(fmt.Sprintf(" %d ", root.Val))
		root = root.Next
	}
	return str.String()
}

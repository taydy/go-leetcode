package easy

import (
	"github.com/taydy/go-leetcode/structure"
	"testing"
)

func TestIsPalindromeList(t *testing.T) {
	head := &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
				Next: &structure.ListNode{
					Val: 2,
					Next: &structure.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	if !IsPalindromeList(head) {
		t.FailNow()
	}

	head = &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
				Next: &structure.ListNode{
					Val: 3,
					Next: &structure.ListNode{
						Val: 2,
						Next: &structure.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	if !IsPalindromeList(head) {
		t.FailNow()
	}

	head = &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
				Next: &structure.ListNode{
					Val: 4,
					Next: &structure.ListNode{
						Val: 2,
						Next: &structure.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	if IsPalindromeList(head) {
		t.FailNow()
	}
}

package easy

import (
	"github.com/taydy/go-leetcode/structure"
	"testing"
)

func TestGetIntersectionNode1(t *testing.T) {
	intersectNode := &structure.ListNode{
		Val: 4,
		Next: &structure.ListNode{
			Val: 5,
			Next: &structure.ListNode{
				Val: 6,
			},
		},
	}
	nodeA := &structure.ListNode{
		Val: 0,
		Next: &structure.ListNode{
			Val:  1,
			Next: intersectNode,
		},
	}
	nodeB := &structure.ListNode{
		Val:  2,
		Next: intersectNode,
	}
	t.Log(GetIntersectionNode(nodeA, nodeB))
}

func TestGetIntersectionNode2(t *testing.T) {
	nodeA := &structure.ListNode{
		Val: 0,
		Next: &structure.ListNode{
			Val: 1,
		},
	}
	nodeB := &structure.ListNode{
		Val: 2,
	}
	t.Log(GetIntersectionNode(nodeA, nodeB))
}

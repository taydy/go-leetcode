package easy

import "testing"

func TestGetIntersectionNode1(t *testing.T) {
	intersectNode := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	nodeA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: intersectNode,
		},
	}
	nodeB := &ListNode{
		Val:  2,
		Next: intersectNode,
	}
	t.Log(GetIntersectionNode(nodeA, nodeB))
}

func TestGetIntersectionNode2(t *testing.T) {
	nodeA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}
	nodeB := &ListNode{
		Val: 2,
	}
	t.Log(GetIntersectionNode(nodeA, nodeB))
}

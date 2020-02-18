package medium

import (
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}

	if views := rightSideView(root); !Equal(views, []int{1, 3, 4}) {
		t.FailNow()
	}

	if views := rightSideView(nil); !Equal(views, []int{}) {
		t.FailNow()
	}
}

package easy

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	exceptDepth := 3
	if maxDepth := MaxDepth(root); maxDepth != exceptDepth {
		t.Logf("except %d, actual %d", exceptDepth, maxDepth)
		t.FailNow()
	}
}

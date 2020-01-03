package medium

import "testing"

func TestLowestCommonAncestor(t *testing.T) {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	if actual := LowestCommonAncestor(root, &TreeNode{
		Val: 5,
	}, &TreeNode{
		Val: 1,
	}); actual == nil || actual.Val != 3 {
		t.Logf("except 3")
		t.FailNow()
	}

	if actual := LowestCommonAncestor(root, &TreeNode{
		Val: 5,
	}, &TreeNode{
		Val: 4,
	}); actual == nil || actual.Val != 5 {
		t.Logf("except 5")
		t.FailNow()
	}

}

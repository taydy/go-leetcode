package easy

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	except := 3
	if result := DiameterOfBinaryTree(root); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}

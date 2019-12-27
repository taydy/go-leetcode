package medium

//	根据一棵树的前序遍历与中序遍历构造二叉树。
//
//	注意:
//		你可以假设树中没有重复的元素。
//
//	例如，给出
//
//	前序遍历 preorder = [3,9,20,15,7]
//	中序遍历 inorder = [9,3,15,20,7]
//	返回如下的二叉树：
//
//	    3
//	   / \
//	  9  20
//	    /  \
//	   15   7
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
//
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	mid := -1
	for i, v := range inorder {
		if v == preorder[0] {
			mid = i
			break
		}
	}

	if mid == -1 {
		return root
	}

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

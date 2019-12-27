package medium

//	根据一棵树的中序遍历与后序遍历构造二叉树。
//
//	注意:
//		你可以假设树中没有重复的元素。
//
//	例如，给出
//
//	中序遍历 inorder = [9,3,15,20,7]
//	后序遍历 postorder = [9,15,7,20,3]
//	返回如下的二叉树：
//
//	    3
//	   / \
//	  9  20
//	    /  \
//	   15   7
//
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
//
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	pLength := len(postorder)

	root := &TreeNode{
		Val: postorder[pLength-1],
	}

	mid := -1
	for i, v := range inorder {
		if v == root.Val {
			mid = i
			break
		}
	}

	if mid == -1 {
		return root
	}

	root.Left = buildTree1(inorder[:mid], postorder[:mid])
	root.Right = buildTree1(inorder[mid+1:], postorder[mid:pLength-1])

	return root
}

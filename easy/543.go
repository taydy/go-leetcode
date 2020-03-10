package easy

import "github.com/taydy/go-leetcode/util"

//	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
//
//	示例 :
//		给定二叉树
//		          1
//		         / \
//		        2   3
//		       / \
//		      4   5
//	返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//	注意：两结点之间的路径长度是以它们之间边的数目表示。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
//
func DiameterOfBinaryTree(root *TreeNode) int {
	maxLength := 1

	diameterOfBinaryTreeDepth(root, &maxLength)
	return maxLength - 1
}

func diameterOfBinaryTreeDepth(root *TreeNode, maxLength *int) int {
	if root == nil {
		return 0
	}
	l, r := diameterOfBinaryTreeDepth(root.Left, maxLength), diameterOfBinaryTreeDepth(root.Right, maxLength)
	*maxLength = util.MaxInt(*maxLength, l+r+1)
	return util.MaxInt(l, r) + 1
}

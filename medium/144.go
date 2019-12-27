package medium

//	给定一个二叉树，返回它的 前序 遍历。
//
// 	示例:
//
//	输入: [1,null,2,3]
// 	  1
//	    \
//	     2
//	    /
//	   3
//
//	输出: [1,2,3]
//
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
//
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

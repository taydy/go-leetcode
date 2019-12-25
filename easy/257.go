package easy

import (
	"strconv"
)

//	给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//	说明: 叶子节点是指没有子节点的节点。
//
//	示例:
//	输入:
//		   1
//		 /   \
//		2     3
//		 \
//		  5
//
//	输出: ["1->2->5", "1->3"]
//
//	解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-paths
//
func BinaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	paths = helper(root, "", paths)
	return paths
}

func helper(root *TreeNode, path string, paths []string) []string {
	if root != nil {
		path += strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil {
			paths = append(paths, path)
		} else {
			path += "->"
			paths = helper(root.Left, path, paths)
			paths = helper(root.Right, path, paths)
			return paths
		}
	}
	return paths
}
